package main

import (
	_ "client/rpc"
	"common/net"
	"common/util"
	"fmt"
	"protos"
	"runtime"
	"slg/const"
	"time"

	"github.com/golang/protobuf/proto"
)

var DEBUG bool

func init() {
	Util.Info()
}

func main() {
	if runtime.GOOS == "windows" {
		//test()
	}
	fmt.Println(">>client start=========================")

	//测试自连
	SERVER := Net.Connect("localhost:8341", func(conn *Net.Conn) {
		conn.CallOut(Const.Login_C, &protos.Login_C{
			OpenId: proto.String("2017015025"),
			// Uid:        proto.Int64(0),
			// ServerName: proto.String("999"),
		})

	}, func(conn *Net.Conn, pid int, data []byte) {
		conn.CallIn(pid, data)
	}, func(conn *Net.Conn) {
		panic("exit")
	})

	//signal
	for {
		time.Sleep(time.Second * 30)
		SERVER.CallOut(Net.Ping, &protos.Ping{})
	}
}
