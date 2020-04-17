package Troop

import (
	"common/sql"

	// "math"
	"protos"
	"sync"
	"time"
)

const (
	TROOP_MAX = 3 //每人最多部队数
	//(1采,2侦,3撤,4驻,5集,6打怪,7打人,8联攻怪,9联攻人,10联防人)
	TROOP_RES     = 1
	TROOP_SCOUT   = 2
	TROOP_BACK    = 3
	TROOP_STATION = 4
	TROOP_MASS    = 5
	TROOP_PVE     = 6
	TROOP_PVP     = 7
	TROOP_PPVE    = 8
	TROOP_PPVP    = 9
	TROOP_DEF     = 10
)

type TroopMap struct {
	sync.RWMutex
	troops map[int64]*protos.Troop
}

var TroopMgr = TroopMap{troops: make(map[int64]*protos.Troop)}

func (m *TroopMap) Set(sid int64, t *protos.Troop) {
	m.Lock()
	if t == nil {
		delete(m.troops, sid)
	} else {
		m.troops[sid] = t
	}
	m.Unlock()
}
func (m *TroopMap) Get(sid int64) *protos.Troop {
	m.Lock()
	defer m.Unlock()
	return m.troops[sid]
}

func init() {
	// row, err := Sql.Query("select * from w_troop")
	// if err != nil {
	// 	panic(err)
	// }
	// for row.Next() {
	// 	var sx, sy, tp, tx, ty int32
	// 	var sid, uid, tuid, st, et int64
	// 	row.Scan(&sid, &uid, &sx, &sy, &tp, &tx, &ty, &tuid, &st, &et)
	// 	t := protos.Troop{Sid: &sid, Uid: &uid, Sx: &sx, Sy: &sy, Tp: &tp, Tx: &tx, Ty: &ty, Tuid: &tuid, St: &st, Et: &et}

	// 	TroopMgr.Set(sid, &t)
	// }
	// go onTickTroops()

}

func onTickTroops() {

	for {
		time.Sleep(time.Second)
		now := (time.Now().UnixNano() / 1e6)
		//sort.Sort(byEt
		for sid, troop := range TroopMgr.troops {
			if troop.GetEt() > 0 && troop.GetEt() < now {
				switch troop.GetTp() {
				}
				Sql.Exec("delete from w_troop where sid=?", sid)

				TroopMgr.Set(sid, nil)
			}
			// 	ss.CallOut( 12, &protos.Response_S{ProtoId: proto.Int32(0),
			// 		Updates: &protos.Updates{},
			// 	})
		}
	}
}

func TroopMarch(uid, sid, tuid int64, tp, tx, ty int32) {
	// ret := Sql.Query2Map1("select cityX,cityY from u_user where uid=?;", uid)
	// sx := int32(ret["cityX"].(int64))
	// sy := int32(ret["cityY"].(int64))
	// tile := Tiles[sy][sx]
	// switch tp {
	// case TROOP_RES:
	// 	if tile.Tp < 1 || tile.Tp > 4 {
	// 		return
	// 	}

	// case TROOP_SCOUT:

	// case TROOP_BACK: //只能对非联攻非返程的外出部队
	// 	if sid == 0 {
	// 		return
	// 	}

	// case TROOP_STATION:

	// case TROOP_MASS:

	// case TROOP_PVE:

	// case TROOP_PVP:
	// case TROOP_PPVE:
	// case TROOP_PPVP:
	// case TROOP_DEF:
	// default:
	// }
	// dis := int64(math.Sqrt(math.Pow(float64(sx)-float64(tx), 2) + math.Pow(float64(sy)-float64(ty), 2)))

	// now := (time.Now().UnixNano() / 1e6)
	// et := now + dis

	// t := protos.Troop{Sid: &sid, Uid: &uid, Sx: &sx, Sy: &sy, Tp: &tp, Tx: &tx, Ty: &ty, Tuid: &tuid, St: &now, Et: &et}
	// _, sid, _ = Sql.Exec("insert into w_troop(uid,sx,sy,tp, tx,ty,tuid,st,et)values(?,?,?,?, ?,?,?,?,?)", uid, sx, sy, tp, tx, ty, tuid, now, et)
	// t.Sid = &sid
	// TroopMgr.Set(sid, &t)
	//ss := getSightByXY(sx, sy)

}

//经过的视区
func getCrossSights(sx, sy, tx, ty int32) {
	var tmp int32
	if sx > tx {
		tmp = sx
		sx = tx
		tx = tmp
	}
	if sy > ty {
		tmp = sy
		sy = ty
		ty = tmp
	}
	//t := []*Sight{}

	// for r := sy / SIGHT_WIDTH; r <= ty/SIGHT_WIDTH; r++ {

	// 	for c := sx / SIGHT_WIDTH; c <= tx/SIGHT_WIDTH; c++ {

	// 	}
	// }
}
