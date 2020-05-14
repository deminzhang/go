package Net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	BUF_MAX_LEN = 1024 * 8 //TODO 最大包长 转到配置文件
)

type Conn struct {
	net.Conn
	uid  int64
	Call func()
}

func (conn *Conn) GetUid() int64 {
	return conn.uid
}
func (conn *Conn) SetUid(uid int64) {
	if uid == 0 {
		SetUid(uid, nil)
	} else {
		SetUid(uid, conn)
	}
	conn.uid = uid

}
func (conn *Conn) ReadLen(lenth int, timeOut time.Duration) ([]byte, error) {
	conn.SetDeadline(time.Now().Add(timeOut))
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
func (conn *Conn) Send(data []byte) {
	plen := len(data)
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
func (conn *Conn) SendRpc(pid int, data []byte) {
	head := make([]byte, 4)
	plen := len(data)

	binary.BigEndian.PutUint32(head, uint32(4+plen))
	conn.Write(head)
	binary.BigEndian.PutUint32(head, uint32(pid)<<16)
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

func (conn *Conn) CallIn(pid int, buf []byte) {
	uid := conn.GetUid()
	rpc := rpcF[pid]
	if rpc == nil {
		log.Println(">>CallIn.Default", pid)
		rpc = rpcF[Response_S]
		if rpc == nil {
			log.Fatal(">>NoRegRpc", Response_S)
		}
	}
	//begin
	rpc(conn, pid, buf, uid)
	//commit
}

func (conn *Conn) CallOut(pid int, pb proto.Message) {
	// buf, err := Encode(pb)
	buf, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}
	conn.SendRpc(pid, buf)
}

func (conn *Conn) Decode(buf []byte, pb proto.Message) bool {
	if err := proto.Unmarshal(buf, pb); err != nil {
		log.Println("Unmarshal error:", err)
		conn.Close()
		return false
	}
	return true
}

//------------------------------------------------------------------------------
//uid->*Conn集
const (
	CONN_MAP_GROUP_NUM = 64
)

type netMap struct {
	sync.RWMutex
	list map[int64]*Conn
}

type netMaps struct {
	groupNum int64
	lists    []*netMap
}

func (g *netMaps) Set(k int64, v *Conn) {
	i := k % g.groupNum
	m := g.lists[i]
	m.Lock()
	if v == nil {
		m.list[k] = v
	} else {
		delete(m.list, k)
	}
	m.Unlock()
}
func (g *netMaps) Get(k int64) *Conn {
	i := k % g.groupNum
	m := g.lists[i]
	m.Lock()
	defer m.Unlock()
	return m.list[k]
}

var AllNets *netMaps

func init() {
	AllNets = &netMaps{}
	for i := 0; i < CONN_MAP_GROUP_NUM; i++ {
		AllNets.lists = append(AllNets.lists, &netMap{list: make(map[int64]*Conn)})
		AllNets.groupNum++
	}
}

func GetByUid(uid int64) *Conn {
	return AllNets.Get(uid)
}

func SetUid(uid int64, conn *Conn) {
	old := AllNets.Get(uid)
	if old == nil {
		AllNets.Set(uid, conn)
	} else {
		//TODO 重连或顶踢
	}
}

//Server------------------------------------------------------------------------
//客户连接
var _connInCount = 0

func onAccept(conn *Conn, onListen func(*Conn), onData func(*Conn, int, []byte), onClose func(*Conn)) {
	fmt.Println("onListen:", conn.RemoteAddr(), conn.LocalAddr())

	defer func() {
		fmt.Println("onClientClose:", conn.RemoteAddr(), conn.LocalAddr())
		_connInCount--
		log.Println(">>clientCount=", _connInCount)
		onClose(conn)
		conn.Close()
	}()
	_connInCount++
	log.Println(">>clientCount=", _connInCount)
	onListen(conn)

	timeOut := time.Second * 10 //首包头超时 小
	for {
		head, err := conn.ReadLen(4, timeOut)
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
		//onHead(conn,head)
		headInt := binary.BigEndian.Uint32(head)
		// pid, plen := headInt>>16, headInt<<16>>16

		timeOut = time.Second * 10 //包体超时 小
		body, err := conn.ReadLen(int(headInt), timeOut)
		if err != nil {
			fmt.Println(err)
			return
		}
		//onBody(conn,body)
		pidb := body[:2]
		pid := int(binary.BigEndian.Uint16(pidb))
		fmt.Printf(">>S:recvLen/Pid(%d/%d)\n", int(headInt), pid)
		onData(conn, pid, body[4:])
		timeOut = time.Minute * 10 //非首次包头超时 大
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
			connEx := &Conn{
				Conn: conn,
			}
			go onAccept(connEx, onListen, onData, onClose)
		}
	}()
	return ln
}

//Client------------------------------------------------------------------------
func onConnect(conn *Conn, onConn func(*Conn), onData func(*Conn, int, []byte), onDisconn func(*Conn)) {
	fmt.Println("onConnect:", conn.RemoteAddr(), conn.LocalAddr())

	onConn(conn)
	defer func() {
		onDisconn(conn)
		fmt.Println("onDisconn:", conn.RemoteAddr())
		conn.Close()
	}()
	for {
		head, err := conn.ReadLen(4, time.Second*30)
		if err != nil {
			fmt.Println(err)
			return
		}
		//onHead(conn,head)
		headInt := binary.BigEndian.Uint32(head)
		body, err := conn.ReadLen(int(headInt), time.Second*10)
		if err != nil {
			fmt.Println(err)
			return
		}
		//onHead(conn,body)
		pidb := body[:2]
		pid := int(binary.BigEndian.Uint16(pidb))
		fmt.Printf(">>C:recvLen/Pid(%d/%d)\n", int(headInt), pid)
		onData(conn, pid, body[4:])
	}
}
func Connect(addr string, onConn func(*Conn), onData func(*Conn, int, []byte), onDisconn func(*Conn)) *Conn {
	fmt.Println(">>Connecting:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
		return nil
	}
	connEx := &Conn{
		Conn: conn,
	}
	go onConnect(connEx, onConn, onData, onDisconn)
	return connEx
}
