package World

import (
	"common/net"

	"log"
	"protos"
	"slg/const"

	"github.com/golang/protobuf/proto"
)

//RPC
func init() {

	Net.RegRpc(Const.View_C, func(c *Net.Conn, pid int, buf []byte, uid int64) {
		ps := protos.View_C{}
		if !c.Decode(buf, &ps) {
			return
		}
		log.Println("<<<View_C", ps.GetServer(), ps.GetX(), ps.GetY())

		// list := moveEyes(0, uid, ps.GetX(), ps.GetY(), ss.Get("sightX"), ss.Get("sightY"))
		// log.Println("<<<sendView_C Num", len(list))
		// //删的前端自理
		updates := &protos.Updates{}
		// for _, o := range list {
		// 	o.AppendTo(updates)
		// }
		c.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(int32(pid)),
			Props: updates,
		})
	})

	Net.RegRpc(Const.CityMove_C, func(c *Net.Conn, pid int, buf []byte, uid int64) {
		ps := protos.CityMove_C{}
		if !c.Decode(buf, &ps) {
			return
		}
		log.Println("<<<CityMove_C", ps.GetServer(), ps.GetX(), ps.GetY())
	})

}
