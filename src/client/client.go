package main

import (
	"common/net"
	"common/util"
	"fmt"
	"protocol"
	"runtime"
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
	go Net.Connect("localhost:8341", func(ss *Net.Session) {
		ss.CallOut(Net.Ping, &protos.Ping{})
		Net.Send(ss.Conn, 1, []byte("ClientPing"))
		ss.CallOut(91, &protos.Login_C{
			OpenId: proto.String("2017015025"),
			// Uid:        proto.Int64(0),
			// ServerName: proto.String("999"),
		})

	}, func(ss *Net.Session) {
		panic("exit")
	})

	//signal
	for {
		time.Sleep(time.Second * 1)
	}
}
