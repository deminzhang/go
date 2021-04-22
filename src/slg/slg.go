package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"common/event"
	"common/sql"
	"common/util"

	"slg/const"

	"slg/config"
	_ "slg/item"
	_ "slg/mail"
	"slg/server"
	_ "slg/test"
	_ "slg/ticker"
	_ "slg/user"
	_ "slg/world"

	// _ "slg/building"
	// _ "slg/job"
	// _ "slg/troop"
	// _ "slg/unit"

	"github.com/BurntSushi/toml"
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

	//开启网络监听
	Listen(conf.Listen)
	//心跳监控
	for {
		time.Sleep(time.Second)
	}
}
