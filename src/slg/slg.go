package main

import (
	"log"
	"protos"
	"runtime"
	"time"

	"common/event"
	"common/net"
	"common/sql"
	"common/util"

	"slg/const"
	_ "slg/user"

	_ "slg/building"
	_ "slg/item"
	_ "slg/job"
	_ "slg/mail"
	_ "slg/server"
	_ "slg/ticker"
	_ "slg/unit"

	_ "slg/troop"
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
	ServerID  int64
}

func init() {
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
	//Cfg.Load("./")
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
	//Server.Init()

	//载入时间管理
	// Ticker.Init()

	log.Println("Event.OnServerStart...")
	Event.Call(Const.OnServerStart)

	//开启网络监听
	go Net.Listen(conf.Listen)

	//测试自连
	//client, server := net.Pipe()
	if runtime.GOOS == "windows" {
		go Net.Connect("localhost:8341", func(ss Net.Session) {
			ss.CallOut(Net.Ping, &protos.Ping{})
			ss.Send(1, []byte("SelfPing"))

			ss.CallOut(Const.Login_C, &protos.Login_C{
				OpenId: proto.String("2017015999622"),
				//Uid:    proto.Int64(0),
			})
		}, func(ss Net.Session) {
			log.Println("onClose", ss)
		})
	}
	//心跳监控
	for {
		time.Sleep(time.Second)
	}
}
