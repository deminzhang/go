package main

import (
	Net "common/net"
	"encoding/binary"
	"log"
	"math/rand"
	"runtime"
	"time"

	Event "common/event"
	Sql "common/sql"

	Const "slg/const"

	Config "slg/config"
	_ "slg/item"
	_ "slg/mail"
	Server "slg/server"
	//_ "slg/test"
	_ "slg/ticker"
	_ "slg/user"
	_ "slg/world"

	"github.com/BurntSushi/toml"
)

// 配置
type GameConf struct {
	Listen    string
	DBType    string
	DBHost    string
	DBHostDev string
	ServerID  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
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

	//开启网络监听
	// Listen(conf.Listen)
	/*server*/
	//开启网络监听
	go Net.Listen(":8341", func(c *Net.Conn) {
		log.Println("S_onListen")
		c.RecvData(4, func(c *Net.Conn, buf []byte) {
			pid := int(binary.BigEndian.Uint16(buf[:2]))
			log.Println("S_onData", pid)
			c.CallIn(pid, buf[4:])
		}, time.Second*10, time.Minute*1, time.Second*10)
	}, func(conn *Net.Conn, msg string) {
		log.Println("S_onClose")
	})

	/*client*/
	//测试自连
	SERVER := Net.Connect("localhost:8341", func(c *Net.Conn) {
		log.Println("C_onConn")
		c.RecvData(4, func(c *Net.Conn, buf []byte) {
			pid := int(binary.BigEndian.Uint16(buf[:2]))
			log.Println("C_onData", pid)
			c.CallIn(pid, buf[4:])
		}, time.Second*10, time.Minute*1, time.Second*10)

	}, func(c *Net.Conn, err interface{}) {
		log.Println("C_onClose", err)
	})
	//心跳监控
	for {
		time.Sleep(time.Second)
	}
	SERVER.Close()
}
