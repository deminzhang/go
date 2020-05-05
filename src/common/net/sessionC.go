package Net

import (
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

//---------------------------------------------------------
//基础SessionClient
type SessionC struct {
	net.Conn
	uid     int64 //userID
	protoId int32 //protoId
	errCode int32 //
	err     error //
	kvmap   map[string]int
}

//外调
func (ss *SessionC) Close() {
	ss.Conn.Close()
}

//外调
func (ss *SessionC) Send(pid int32, data []byte) {
	Send(ss.Conn, pid, data)
}

//外调
func (ss *SessionC) CallOut(pid int32, msg proto.Message) {
	conn := ss.Conn
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	Send(conn, int32(pid), data)
}

//内调
func (ss *SessionC) CallIn(pid int32, data []byte) {
	//fmt.Println(">>Session.CallIn", pid, data)
	ss.protoId = pid
	uid := ss.GetUid()
	defer ss.afterCall()
	rpc := rpcs[pid]
	if rpc == nil {
		log.Println(">>CallIn.Default", pid, uid)
		rpc = rpcs[Response_S]
		if rpc == nil {
			log.Fatal(">>NoRegRpc", Response_S)
		}
	}
	var sss Session
	sss = ss
	rpc(sss, pid, data, uid)
}

//在线设置
func (ss *SessionC) SetUid(uid int64) {
	if uid == 0 {
		G_uid2session.Set(ss.GetUid(), nil)
	} else {
		G_uid2session.Set(uid, ss)
	}
	ss.uid = uid

}
func (ss *SessionC) GetUid() int64 {
	return ss.uid
}
func (ss *SessionC) GetProtoId() int32 {
	return ss.protoId
}

func (ss *SessionC) onError() {
	log.Println(ss.err)
}

func (ss *SessionC) commit() {

}

func (ss *SessionC) afterCall() {
	ss.errCode = 0
	ss.err = nil
	ss.protoId = 0
}

//错误自动断开
func (ss *SessionC) Error(err error) bool {
	if err == nil {
		return true
	}
	ss.err = err
	ss.onError()
	ss.Close()
	return false
}

//错误自动断开
func (ss *SessionC) Assert(b bool, err error) bool {
	if !b {
		ss.Error(err)
	}
	return b
}

//自定义数值
func (ss *SessionC) Get(key string) int {
	return ss.kvmap[key]
}
func (ss *SessionC) Set(key string, val int) {
	ss.kvmap[key] = val
}

//解码错误自动断开
func (ss *SessionC) Decode(buf []byte, pb proto.Message) bool {
	err := proto.Unmarshal(buf, pb)
	return ss.Error(err)
}
