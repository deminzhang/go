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

	Net.RegRPC(Const.View_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.View_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		log.Println("<<<View_C", ps.GetServer(), ps.GetX(), ps.GetY())

		list := moveEyes(0, uid, ps.GetX(), ps.GetY())
		log.Println("<<<sendView_C Num", len(list))
		//删的前端自理
		updates := &protos.Updates{}
		for _, o := range list {
			o.AppendTo(updates)
		}
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})

	Net.RegRPC(Const.CityMove_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.CityMove_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		log.Println("<<<CityMove_C", ps.GetServer(), ps.GetX(), ps.GetY())
		// x, y := ps.GetX(), ps.GetY()
		// //TODO Item.CheckCost
		// // if !CheckCityLand(x, y) {
		// // 	ss.PostError(protoId, 2, "CityMove_C.noMoveCity")
		// // 	return
		// // }
		// Item.Del(uid, 2, 1, "CityMove")
		// // MoveCity(uid, x, y)

		// Sql.Exec("update u_user set cityX=?,cityY=? where uid=?", x, y, uid)
		// uu := &protos.User{
		// 	Uid:   proto.Int64(uid),
		// 	CityX: proto.Int32(x),
		// 	CityY: proto.Int32(y),
		// }
		// updates := &protos.Updates{}
		// updates.User = uu
		// ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
		// 	Props: updates,
		// })
	})

}
