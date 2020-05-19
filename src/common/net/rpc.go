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

var rpcF = make(map[int]func(*Conn, int, []byte, int64))

// var rpcD = make(map[int]proto.Message)

func init() {
	RegRpc(Ping, func(conn *Conn, pid int, data []byte, uid int64) {
		fmt.Println("<<<Ping", data)
		conn.SendRpc(Pong, data)
	})
	RegRpc(Pong, func(conn *Conn, pid int, data []byte, uid int64) {
		fmt.Println(">>>Pong", data)
	})
}

// func DefRpc(pid int, pb proto.Message) {
// 	if rpcD[pid] != nil {
// 		log.Fatalf("RegRpcC duplicated %d", pid)
// 	}
// 	rpcD[pid] = pb
// }

func RegRpc(pid int, call func(*Conn, int, []byte, int64)) {
	if rpcF[pid] != nil {
		log.Fatalf("RegRpcC duplicated %d", pid)
	}
	rpcF[pid] = call
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
