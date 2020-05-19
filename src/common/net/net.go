package Net

import (
	"strconv"
	// "bytes"
	"encoding/binary"
	"errors"
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
	uid     int64
	OnError func(*Conn, error)
	err     error
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
func (conn *Conn) ReadLen(lenth int, timeOut time.Duration, onData func(*Conn, []byte)) {
	conn.SetDeadline(time.Now().Add(timeOut))
	if lenth == 0 {
		return
	}
	data := make([]byte, lenth)
	l := 0 //长包等合并
	for {
		n, err := conn.Read(data[l:])
		if err != nil {
			conn.err = err
			conn.Close()
			return
		}
		l += n
		if l >= lenth {
			break
		}
	}
	onData(conn, data)
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
			conn.err = errors.New("NoRegRpc pid:" + strconv.Itoa(Response_S))
			conn.Close()
			return
		}
	}
	//begin
	rpc(conn, pid, buf, uid)
	//commit
}

func (conn *Conn) CallOut(pid int, pb proto.Message) {
	buf, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}
	conn.SendRpc(pid, buf)
}
func (conn *Conn) Call(pid int, pb interface{}) {
	buf, err := proto.Marshal(pb.(proto.Message))
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

func onAccept(conn *Conn, onListen func(*Conn), onClose func(*Conn, error)) {
	fmt.Println("onListen:", conn.RemoteAddr(), conn.LocalAddr())
	defer func() {
		fmt.Println("onClientClose:", conn.RemoteAddr(), conn.LocalAddr())
		onClose(conn, conn.err)
		conn.Close()
	}()
	onListen(conn)
}

func Listen(addr string, onListen func(*Conn), onClose func(*Conn, error)) net.Listener {
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
				// continue
				return
			}
			connEx := &Conn{
				Conn:    conn,
				OnError: onClose,
			}
			go onAccept(connEx, onListen, onClose)
		}
	}()
	return ln
}
func ListenUnix() {

}

//Client------------------------------------------------------------------------
func onConnect(conn *Conn, onConn func(*Conn), onDisconn func(*Conn, error)) {
	fmt.Println("onConnect:", conn.RemoteAddr(), conn.LocalAddr())
	defer func() {
		onDisconn(conn, conn.err)
		fmt.Println("onDisconn:", conn.RemoteAddr())
		conn.Close()
	}()
	onConn(conn)
}
func Connect(addr string, onConn func(*Conn), onDisconn func(*Conn, error)) *Conn {
	fmt.Println(">>Connecting:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
		return nil
	}
	connEx := &Conn{
		Conn:    conn,
		OnError: onDisconn,
	}
	go onConnect(connEx, onConn, onDisconn)
	return connEx
}
