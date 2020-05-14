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
	Net.RegRpc(Const.March_C, func(c *Net.Conn, pid int, buf []byte, uid int64) {
		ps := protos.March_C{}
		if !c.Decode(buf, &ps) {
			return
		}
		log.Println("<<<March_C", ps.GetTp(), ps.GetX(), ps.GetY())

		//TroopMarch(uid, ps.GetSid(), 0, 1, ps.GetX(), ps.GetY())

	})

}
