package main

import (
	"common/net"
	"common/sql"
	"common/util"
	"fmt"
	"log"
	"protocol"
	"runtime"
	"time"

	// _ "slg/city"
	// _ "slg/item"
	// _ "slg/mail"
	// _ "slg/rpc"
	// _ "slg/server"
	// _ "slg/user"
	// _ "slg/world"

	_ "github.com/golang/protobuf/proto"
)

func init() {
	Util.Info()
}

func main() {
	fmt.Println(">>server start=========================")
	//载入配置
	//Cfg.Load("./")
	//数据库初始化与更新
	Sql.Init()
	//服务器数据
	//Server.Init()
	//载入时间管理
	//Ticker.Init()
	//开启网络监听
	go Net.Listen(":8341")

	//测试自连
	//client, server := net.Pipe()
	if runtime.GOOS == "windows" {
		//go Net.Connect("localhost:8341", func(conn net.Conn) {
		go Net.Connect("localhost:8341", func(ss *Net.Session) {
			ss.CallOut(Net.Ping, &protos.Ping{})
			Net.Send(ss.Conn, 1, []byte("LoginPing"))
			// ss.CallOut(91, &protos.Login_C{
			// 	OpenId: proto.String("2017015939"),
			// 	//Uid:    proto.Int64(0),
			// })

		}, func(ss *Net.Session) {
			log.Println("onClose", ss)
		})
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
