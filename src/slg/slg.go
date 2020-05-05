package main

import (
	"log"
	"math/rand"
	"protos"
	"runtime"
	"time"

	"common/event"
	"common/net"
	"common/sql"
	"common/util"

	"slg/const"

	"slg/config"
	"slg/server"
	_ "slg/user"

	_ "slg/building"
	_ "slg/item"
	_ "slg/job"
	_ "slg/mail"
	_ "slg/test"
	_ "slg/ticker"
	_ "slg/troop"
	_ "slg/unit"
	_ "slg/world"

	"github.com/BurntSushi/toml"
	"github.com/golang/protobuf/proto"
)

//配置
type GameConf struct {
	Listen    string
	DBType    string
	DBHost    string
	DBHostDev string
	ServerID  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
	Util.Info()
}

func main() {
	log.Println(">>server start=========================")
	defer log.Println(">>server error exit=========================")
	//载入系统配置
	var conf GameConf
	if _, err := toml.DecodeFile("./game.toml", &conf); err != nil {
		log.Fatal(err)
	}
	log.Println("conf:", conf)

	//载入数值配置
	Config.Load("./")
	log.Println("Event.OnLoadConfig...")
	Event.Call(Const.OnLoadConfig)
	Event.Call(Const.OnCheckConfig)

	//数据库初始化与更新
	if runtime.GOOS == "windows" {
		Sql.ORMConnect(conf.DBType, conf.DBHostDev)
	} else {
		Sql.ORMConnect(conf.DBType, conf.DBHost)
	}
	log.Println("Event.OnInitDB...")
	Event.Call(Const.OnInitDB)
	Event.Call(Const.OnLoadDB)

	//服务器数据
	Server.Init(conf.ServerID)

	//载入时间管理
	// Ticker.Init()

	log.Println("Event.OnServerStart...")
	Event.Call(Const.OnServerStart)

	Net.InitSession(64)
	//开启网络监听
	Net.Listen(conf.Listen, func(conn *Net.Conn) {
		log.Println("onListen")
		conn.Session = &Net.SessionS{Conn: conn.Conn}
	}, func(conn *Net.Conn, pid int, data []byte) {
		// log.Println("onData", pid)
		Net.CallIn(conn, pid, data)
	}, func(conn *Net.Conn) {
		ss := conn.Session
		if ss != nil && ss.GetUid() > 0 {
			Event.Call(Const.OnUserOffline, ss.GetUid())
		}
	})

	//测试自连
	//client, server := net.Pipe()
	if runtime.GOOS == "windows" {

		Net.RegRPC(Const.Response_S, func(ss Net.Session, pid int32, data []byte, uid int64) {
			ps := protos.Response_S{}
			if !ss.Decode(data, &ps) {
				return
			}
			log.Printf("recv<<Response_S srcPid=%d msg=\n", ps.GetProtoId())
			log.Println(ps, ps.GetUpdates())
		})
		Net.RegRPC(Const.Error_S, func(ss Net.Session, pid int32, data []byte, uid int64) {
			ps := protos.Error_S{}
			if !ss.Decode(data, &ps) {
				return
			}
			log.Println("<<<Error_S", pid, ps.GetCode(), ps.GetMsg())
		})
		Net.RegRPC(Const.Login_S, func(ss Net.Session, pid int32, data []byte, uid int64) {
			ps := protos.Login_S{}
			if !ss.Decode(data, &ps) {
				return
			}
			log.Println("recv>>Login_S:\n", Const.Login_S, len(data), ps)
			ss.CallOut(Const.GetRoleInfo_C, &protos.GetRoleInfo_C{})
		})
		Net.RegRPC(Const.GetRoleInfo_S, func(ss Net.Session, pid int32, data []byte, uid int64) {
			log.Println("recv>>GetRoleInfo_S\n", Const.GetRoleInfo_S, len(data))
			ss.CallOut(Const.View_C, &protos.View_C{
				X: proto.Int32(16),
				Y: proto.Int32(16),
			})

		})

		go Net.Connect("localhost:8341", func(conn *Net.Conn) {
			conn.Session = &Net.SessionC{Conn: conn.Conn}
			Net.Send(conn, Net.Ping, []byte("SelfPing"))

			Net.CallOut(conn, Const.Login_C, &protos.Login_C{
				OpenId: proto.String("20170159c326"),
				//Uid:    proto.Int64(0),
			})
		}, func(conn *Net.Conn, pid int, data []byte) {
			// log.Println("onData", pid)
			Net.CallIn(conn, pid, data)
		}, func(conn *Net.Conn) {
			log.Println("onClose", conn.RemoteAddr())
		})
	}
	//心跳监控
	for {
		time.Sleep(time.Second)
	}
}
