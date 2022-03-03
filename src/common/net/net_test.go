package net_test

import (
	Net "common/event"
	"testing"
)
func TestNet(t *testing.T) {
	ExampleServer()
	ExampleClient()
	// for {
	// 	time.Sleep(time.Second)
	// }
}

func ExampleServer()
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
}
func ExampleClient(){
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
}