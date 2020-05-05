package Net

import (
	"bytes"
	// "common/event"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	BUF_MAX_LEN = 1024 * 8 //TODO 最大包长 转到配置文件
)

type Conn struct {
	net.Conn
	Session Session
}

//TODO 多协程读写map需加锁 sync.Mutex or sync.RwMutex
//客户连接
var _clientCount = 0

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

func ReadLen(conn net.Conn, lenth int) ([]byte, error) {
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

//Server------------------------------------------------------------------------
func onAccept(conn net.Conn, onListen func(*Conn), onData func(*Conn, int, []byte), onClose func(*Conn)) {
	fmt.Println("onListen:", conn.RemoteAddr(), conn.LocalAddr())
	conn.SetDeadline(time.Now().Add(time.Second * 30)) //初次收不到包超时

	connEx := &Conn{
		Conn: conn,
	}
	onListen(connEx)
	_clientCount++
	log.Println(">>clientCount=", _clientCount)
	defer func() {
		fmt.Println("onNetClose:", conn.RemoteAddr(), conn.LocalAddr())
		_clientCount--
		log.Println(">>clientCount=", _clientCount)
		onClose(connEx)
		conn.Close()
	}()
	for {
		head, err := ReadLen(conn, 4)
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
		data, err := ReadLen(conn, int(plen))
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.SetDeadline(time.Now().Add(time.Minute * 10)) //以后的包超时
		onData(connEx, int(pid), data)
	}
}

//func Listen(addr string, onListen, onClose) {
func Listen(addr string, onListen func(*Conn), onData func(*Conn, int, []byte), onClose func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Print("Listen.err", err)
		return nil
	}
	fmt.Println(">>listening:", ln.Addr())
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Print(err)
				continue
			}
			go onAccept(conn, onListen, onData, onClose)
		}
	}()
	return ln
}

//Client------------------------------------------------------------------------
func onConnect(conn net.Conn, onConn func(*Conn), onData func(*Conn, int, []byte), onDisconn func(*Conn)) {
	fmt.Println("onConnect:", conn.RemoteAddr(), conn.LocalAddr())

	connEx := &Conn{
		Conn: conn,
	}
	onConn(connEx)
	defer func() {
		onDisconn(connEx)
		fmt.Println("onDisconn:", conn.RemoteAddr(), conn.LocalAddr())
		conn.Close()
	}()
	for {
		head, err := ReadLen(conn, 4)
		if err != nil {
			fmt.Println(err)
			return
		}
		headInt := binary.BigEndian.Uint32(head)
		pid, plen := headInt>>16, headInt<<16>>16
		//fmt.Printf(">>recv(%d:%d)\n", pid, plen)
		//fmt.Printf(">>recvX(%x:%x)\n", pid, plen)
		data, err := ReadLen(conn, int(plen))
		if err != nil {
			fmt.Println(err)
			return
		}
		onData(connEx, int(pid), data)
	}
}
func Connect(addr string, onConn func(*Conn), onData func(*Conn, int, []byte), onDisconn func(*Conn)) net.Conn {
	fmt.Println(">>Connecting:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
		return nil
	}
	go onConnect(conn, onConn, onData, onDisconn)
	return conn
}
