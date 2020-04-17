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
	Net.RegRPC(Const.March_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.March_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		log.Println("<<<March_C", ps.GetTp(), ps.GetX(), ps.GetY())

		//TroopMarch(uid, ps.GetSid(), 0, 1, ps.GetX(), ps.GetY())

	})

}
