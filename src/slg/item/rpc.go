package Item

import (
	"common/net"
	"log"
	"protos"
	"slg/const"

	"github.com/golang/protobuf/proto"
)

//RPC
func init() {

	Net.RegRpc(Const.ItemUse_C, func(c *Net.Conn, pid int, buf []byte, uid int64) {
		log.Println(">>>ItemUse_C", buf)
		ps := protos.ItemUse_C{}
		if err := proto.Unmarshal(buf, &ps); err != nil {
			log.Println("Decode error: ", err, buf)
			c.Close()
			return
		}
		log.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Use(uid, ps.GetCid(), int64(ps.GetNum()))

	})
	Net.RegRpc(Const.ItemDel_C, func(c *Net.Conn, pid int, buf []byte, uid int64) {
		log.Println(">>>ItemUse_C", buf)
		ps := protos.ItemDel_C{}
		if err := proto.Unmarshal(buf, &ps); err != nil {
			log.Println("Decode error: ", err, buf)
			c.Close()
			return
		}
		log.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Del(uid, ps.GetCid(), int64(ps.GetNum()), "del")

	})

}
