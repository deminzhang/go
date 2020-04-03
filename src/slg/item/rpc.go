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

	Net.RegRPC(Const.ItemUse_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		log.Println(">>>ItemUse_C", data)

		ps := protos.ItemUse_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("View_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		log.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Use(uid, ps.GetCid(), int64(ps.GetNum()))

	})
	Net.RegRPC(Const.ItemDel_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		log.Println(">>>ItemUse_C", data)

		ps := protos.ItemDel_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("View_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		log.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Del(uid, ps.GetCid(), int64(ps.GetNum()), "del")

	})

}
