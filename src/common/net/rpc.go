package Net

import (
	"fmt"
	"log"
	"protocol"

	"github.com/golang/protobuf/proto"
)

const (
	Ping       = 1
	Pong       = 2
	Response_S = 12
	Error_S    = 14
)

//TODO 多协程读写map需加锁 sync.Mutex or sync.RwMutex
var rpcs = make(map[int32]func(*Session, int32, int64, []byte))
var decoders = make(map[int32]interface{})

func init() {
	//go tickServer()
	RegRPC(Ping, func(ss *Session, pid int32, uid int64, data []byte) {
		fmt.Println("<<<Ping", data)
		Send(ss.Conn, Pong, data)
	})
	RegRPC(Pong, func(ss *Session, pid int32, uid int64, data []byte) {
		fmt.Println(">>>Pong", data)
	})

}

func RegRPC(pid int32, call func(*Session, int32, int64, []byte)) {
	if rpcs[pid] != nil {
		log.Fatalf("RegRPC duplicated %d", pid)
	}
	rpcs[pid] = call
	//decoders[pid]=?
}

func ConnByUid(uid int64) *Session {
	ss := G_uid2session.Get(uid)
	if ss == nil {
		return nil
	}
	return ss
}

func CallUid(uid int64, pid int32, msg proto.Message) {
	ss := G_uid2session.Get(uid)
	if ss == nil {
		return
	}
	ss.CallOut(pid, msg)
}

func CallUids(uids []int64, pid int32, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for _, uid := range uids {
		ss := G_uid2session.Get(uid)
		if ss != nil {
			Send(ss.Conn, pid, data)
		}
	}
}

func CallUidError(uid int64, pid int32, code int32, msg string) {
	ss := G_uid2session.Get(uid)
	if ss == nil {
		return
	}
	ss.CallOut(Error_S, &protos.Error_S{
		ProtoId: proto.Int32(int32(pid)),
		Code:    proto.Int32(code),
		Msg:     proto.String(msg),
	})
}
