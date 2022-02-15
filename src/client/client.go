package main

import (
	_ "client/rpc"
	"common/net"
	"common/utilX"
	"fmt"
	"protos"
	"runtime"
	"slg/const"
	"time"

	"github.com/golang/protobuf/proto"
)

var DEBUG bool

func init() {
	util.Info()
}

func main() {
	if runtime.GOOS == "windows" {
		//test()
	}
	fmt.Println(">>client start=========================")

	//测试自连
	SERVER := Net.Connect("localhost:8341", func(c *Net.Conn) {
		conn.CallOut(Const.Login_C, &protos.Login_C{
			OpenId: proto.String("2017015025"),
			// Uid:        proto.Int64(0),
			// ServerName: proto.String("999"),
		})
		c.RecvData(4, func(c *Net.Conn, buf []byte) {
			pid := int(binary.BigEndian.Uint16(buf[:2]))
			c.CallIn(pid, buf[4:])
		}, time.Second*10, time.Minute*1, time.Second*10)

	}, func(c *Net.Conn, err interface{}) {
		panic(err)
	})

	//signal
	for {
		time.Sleep(time.Second * 30)
		SERVER.CallOut(Net.Ping, &protos.Ping{})
	}
}
