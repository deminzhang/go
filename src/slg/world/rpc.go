package World

import (
	// "common/event"
	"common/net"
	"common/sql"

	"log"
	// "fmt"
	// "math"
	// "math/rand"
	"protos"
	"slg/const"

	"slg/item"
	// "slg/rpc"
	// "sync"
	// "time"
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

		tiles := ResetSight(ss, ps.GetServer(), ps.GetX(), ps.GetY())
		//删的前端自理
		ss.Update().Tile = tiles
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: &protos.Updates{
				Tile: tiles,
			},
		})
	})

	Net.RegRPC(Const.CityMove_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.CityMove_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		log.Println("<<<CityMove_C", ps.GetServer(), ps.GetX(), ps.GetY())
		x, y := ps.GetX(), ps.GetY()
		//TODO Item.CheckCost
		if !CheckCityLand(x, y) {
			ss.PError(protoId, 2, "CityMove_C.noMoveCity")
			return
		}
		Item.Del(uid, 2, 1, "CityMove")
		MoveCity(uid, x, y)

		Sql.Exec("update u_user set cityX=?,cityY=? where uid=?", x, y, uid)
		uu := &protos.User{
			Uid:   proto.Int64(uid),
			CityX: proto.Int32(x),
			CityY: proto.Int32(y),
		}
		updates := &protos.Updates{}
		updates.User = uu
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})

}
