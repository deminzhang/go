package Net

import (
	"fmt"
	"log"
	"net"
	"protos"

	"github.com/golang/protobuf/proto"
)

const (
	Ping       = 1
	Pong       = 2
	Response_S = 12
	Error_S    = 14
)

//TODO 多协程读写map需加锁 sync.Mutex or sync.RwMutex
var rpcs = make(map[int32]func(Session, int32, []byte, int64))
var decoders = make(map[int]interface{})

func init() {
	RegRPC(Ping, func(ss Session, pid int32, data []byte, uid int64) {
		fmt.Println("<<<Ping", data)
		ss.Send(Pong, data)
	})
	RegRPC(Pong, func(ss Session, pid int32, data []byte, uid int64) {
		fmt.Println(">>>Pong", data)
	})
}

func RegRPC(pid int32, call func(Session, int32, []byte, int64)) {
	if rpcs[pid] != nil {
		log.Fatalf("RegRPC duplicated %d", pid)
	}
	rpcs[pid] = call
}

func CallIn(conn *Conn, protoID int, data []byte) {
	if conn.Session != nil {
		conn.Session.CallIn(int32(protoID), data)
	} else {
	}
}

func CallOut(conn net.Conn, protoID int, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	Send(conn, int32(protoID), data)
}

func CallUid(uid int64, protoID int32, msg proto.Message) {
	ss := G_uid2session.Get(uid)
	if ss == nil {
		return
	}
	ss.CallOut(protoID, msg)
}

func CallUids(uids []int64, pid int32, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	for _, uid := range uids {
		ss := G_uid2session.Get(uid)
		if ss != nil {
			ss.Send(pid, data)
		}
	}
}

func CallError(uid int64, pid int32, code int32, msg string) {
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
