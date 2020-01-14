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
	// go Net.Connect("10.45.11.29:18091", func(ss *Net.Session) { //29旧版
	// go Net.Connect("10.45.11.29:8088", func(ss *Net.Session) { //29新版
	// go Net.Connect("10.5.30.213:18091", func(ss *Net.Session) { //?
	// go Net.Connect("10.249.249.156:8088", func(ss *Net.Session) { //FX 10.249.249.156
	go Net.Connect("localhost:8088", func(ss *Net.Session) { //10.249.249.169
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
