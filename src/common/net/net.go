package Net

import (
	"bytes"
	"common/event"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	BUF_MAX_LEN = 1024 * 8 //TODO 最大包长 转到配置文件
)

//TODO 多协程读写map需加锁 sync.Mutex or sync.RwMutex
//客户连接
var _clientCount = 0

//对外连接
//var G_servers = make(map[net.Conn]bool) //需锁

func init() {
	//TODO 读配置文件
	//go tickServer()
}

func Send(conn net.Conn, pid int32, data []byte) {
	var head = make([]byte, 4)
	var plen = int32(len(data))

	binary.BigEndian.PutUint32(head, uint32(pid<<16+plen))
	//fmt.Printf(">>send(%d:%d)\n", pid, plen)

	conn.Write(head)
	if plen == 0 {
		return
	}
	//大包分包
	for plen > BUF_MAX_LEN {
		conn.Write(data[:BUF_MAX_LEN])
		data = data[BUF_MAX_LEN:]
		plen -= BUF_MAX_LEN
		if plen == 0 {
			return
		}
	}
	conn.Write(data)
}

func tickServer() {
	//Test测试
	//fmt.Println(">>ticking")
	for {
		time.Sleep(time.Second * 10)
		//for conn := range G_clients {
		//	Send(conn, 0, []byte("hi,client"))
		//}
	}
}

func readLen(conn net.Conn, lenth int) ([]byte, error) {
	if lenth == 0 {
		return nil, nil
	}
	data := make([]byte, lenth)
	l := 0 //长包等合并
	for l < lenth {
		n, err := conn.Read(data[l:])
		if err != nil {
			return nil, err
		}
		l += n
	}
	return data, nil
}

//Server.listen
func onListen(conn net.Conn) {
	fmt.Println("onListen:", conn.RemoteAddr(), conn.LocalAddr())
	conn.SetDeadline(time.Now().Add(time.Second * 30)) //初次收不到包超时
	_clientCount++
	log.Println(">>clientCount=", _clientCount)
	session := Session{Conn: conn}
	defer func() {
		fmt.Println("onNetClose:", conn.RemoteAddr(), conn.LocalAddr())
		if session.Uid > 0 {
			Event.CallA("OnDisconn", session.Uid)
		}
		_clientCount--
		log.Println(">>clientCount=", _clientCount)
		conn.Close()
	}()
	for {
		head, err := readLen(conn, 4)
		if err != nil {
			fmt.Println(err)
			return
		}
		if bytes.Equal(head, []byte("POST")) {
			fmt.Println("TODO http.POST")
		}
		if bytes.Equal(head, []byte("GET ")) {
			fmt.Println("TODO http.GET ")
		}

		headInt := binary.BigEndian.Uint32(head)
		pid, plen := headInt>>16, headInt<<16>>16
		//fmt.Printf("<<recv(%d:%d)\n", pid, plen)
		//fmt.Printf("<<recvX(%x:%x)\n", pid, plen)

		conn.SetDeadline(time.Now().Add(time.Second * 10)) //收到包头后超时
		data, err := readLen(conn, int(plen))
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.SetDeadline(time.Now().Add(time.Minute * 10)) //以后的包超时
		session.CallIn(int32(pid), data)
		//CallIn(int32(pid), data)
	}
}

//func Listen(addr string, onListen, onClose) {
func Listen(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Print("Listen.err", err)
		return
	}
	fmt.Println(">>listening:", ln.Addr())
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go onListen(conn)
	}
}

//Client.connect
func onConnect(conn net.Conn, onConn func(*Session), onDisconn func(*Session)) {
	fmt.Println("onConnect:", conn.RemoteAddr(), conn.LocalAddr())
	session := Session{Conn: conn}
	onConn(&session)
	defer func() {
		onDisconn(&session)
		fmt.Println("onDisconn:", conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}()
	for {
		head, err := readLen(conn, 4)
		if err != nil {
			fmt.Println(err)
			return
		}
		headInt := binary.BigEndian.Uint32(head)
		pid, plen := headInt>>16, headInt<<16>>16
		//fmt.Printf(">>recv(%d:%d)\n", pid, plen)
		//fmt.Printf(">>recvX(%x:%x)\n", pid, plen)
		data, err := readLen(conn, int(plen))
		if err != nil {
			fmt.Println(err)
			return
		}
		session.CallIn(int32(pid), data)
	}
}
func Connect(addr string, onConn func(*Session), onDisconn func(*Session)) net.Conn {
	fmt.Println(">>Connecting:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
		return nil
	}
	go onConnect(conn, onConn, onDisconn)
	return conn
}
