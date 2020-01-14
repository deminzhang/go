package main

import (
	"common/net"
	"fmt"
	"protocol"
	"slg/rpc"

	_ "github.com/golang/protobuf/proto"
)

func init() {

	Net.RegRPC(Rpc.Response_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Response_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println(">>>Response_S", ps.GetProtoId(), pid, uid, ps)
	})
	Net.RegRPC(Rpc.Error_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Error_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Error_S", ps.GetProtoId(), ps.GetCode(), ps.GetMsg())
	})
	Net.RegRPC(Rpc.Login_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Login_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println(">>>Login_S:", Rpc.Login_S, len(data), ps)

		ss.CallOut(Rpc.GetRoleInfo_C, &protos.GetRoleInfo_C{})

	})

}
