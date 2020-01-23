package main

import (
	"common/net"
	"fmt"
	"protos"
	"slg/const"

	_ "github.com/golang/protobuf/proto"
)

func init() {

	Net.RegRPC(Const.Response_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Response_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println(">>>Response_S", ps.GetProtoId(), pid, uid, ps)
	})
	Net.RegRPC(Const.Error_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Error_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Error_S", ps.GetProtoId(), ps.GetCode(), ps.GetMsg())
	})
	Net.RegRPC(Const.Login_S, func(ss *Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Login_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println(">>>Login_S:", Const.Login_S, len(data), ps)

		ss.CallOut(Const.GetRoleInfo_C, &protos.GetRoleInfo_C{})

	})

}
