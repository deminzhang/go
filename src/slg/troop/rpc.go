package Troop

import (
	"common/net"
	"log"
	"protos"
	"slg/const"
	// "slg/entity"
	// "github.com/golang/protobuf/proto"
)

//rpc
func init() {
	Net.RegRPC(Const.March_C, func(ss Net.Session, pid int32, data []byte, uid int64) {
		ps := protos.March_C{}
		if !ss.Decode(data, &ps) {
			return
		}
		log.Println("<<<March_C", ps.GetTp(), ps.GetX(), ps.GetY())

		//TroopMarch(uid, ps.GetSid(), 0, 1, ps.GetX(), ps.GetY())

	})

}
