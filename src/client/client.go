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
	SERVER := Net.Connect("localhost:8341", func(c *Net.Conn) {
		conn.CallOut(Const.Login_C, &protos.Login_C{
			OpenId: proto.String("2017015025"),
			// Uid:        proto.Int64(0),
			// ServerName: proto.String("999"),
		})

		var onHead, onBody func(*Net.Conn, []byte)
		onHead = func(c *Net.Conn, buf []byte) {
			headInt := int(binary.BigEndian.Uint32(buf))
			c.ReadLen(headInt, time.Second*10, onBody) //包体超时 小
		}
		onBody = func(c *Net.Conn, buf []byte) {
			pid := int(binary.BigEndian.Uint16(buf[:2]))
			c.CallIn(pid, buf[4:])
			c.ReadLen(4, time.Minute*10, onHead)
		}
		c.ReadLen(4, time.Second*10, onHead)
	},  func(c *Net.Conn,, err string) {
		panic(err)
	})

	//signal
	for {
		time.Sleep(time.Second * 30)
		SERVER.CallOut(Net.Ping, &protos.Ping{})
	}
}
