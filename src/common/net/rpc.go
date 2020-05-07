package Net

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

const (
	Ping       = 1
	Pong       = 2
	Response_S = 12
	Error_S    = 14
)

var rpcC = make(map[int]func(*Conn, int, []byte, int64))

// var Decode func(buf []byte, msg interface{}) error
// var Encode func(msg interface{}) ([]byte, error)

func init() {
	// // base
	// Decode = func(buf []byte, msg interface{}) error {
	// 	panic("Net.Decoder Need Rewrite as func(buf []byte, msg interface{}) error")
	// 	return nil
	// }
	// Encode = func(msg interface{}) ([]byte, error) {
	// 	panic("Net.Encoder Need Rewrite as func(msg interface{}) ([]byte, error)")
	// 	return nil, nil
	// }
	// //rewrite
	// Encode = func(msg interface{}) ([]byte, error) {
	// 	return proto.Marshal(msg.(proto.Message))
	// }
	// Decode = func(buf []byte, msg interface{}) error {
	// 	return proto.Unmarshal(buf, msg.(proto.Message))
	// }

	RegRpcC(Ping, func(conn *Conn, pid int, data []byte, uid int64) {
		fmt.Println("<<<Ping", data)
		conn.Send(Pong, data)
	})
	RegRpcC(Pong, func(conn *Conn, pid int, data []byte, uid int64) {
		fmt.Println(">>>Pong", data)
	})
}

func RegRpcC(pid int, call func(*Conn, int, []byte, int64)) {
	if rpcC[pid] != nil {
		log.Fatalf("RegRpcC duplicated %d", pid)
	}
	rpcC[pid] = call
}

func CallUid(uid int64, pid int, pb proto.Message) {
	conn := GetByUid(uid)
	if conn != nil {
		conn.CallOut(pid, pb)
	}
}

func CallError(uid int64, pid int, code int, msg string) {
	conn := GetByUid(uid)
	if conn != nil {

	}
}
