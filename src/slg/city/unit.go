package City

import (
	"common/event"
	"common/net"
	"common/sql"
	"fmt"
	"log"

	//	"net"
	"protos"
	"slg/item"
	"slg/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

func unitUp(uid int64, tp int32, lv int32, num int32) {

}

func ReadUnit(uid int64, tp int32, lv int32) *protos.Unit {
	rows, err := Sql.Query("select num from u_unit where uid=? and tp=? and lv=?", uid)
	if err != nil {
		log.Println("City.OnUserInit error: ", err)
		return nil
	}
	for rows.Next() {
		var num int32
		rows.Scan(&num)
		rows.Close()
		return &protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}
	}
	return nil
}

func AddUnit(uid int64, tp int32, lv int32, num int32) {
	if num == 0 {
		return
	}
	u := ReadUnit(uid, tp, lv)
	if u == nil {
		Sql.Exec("replace into u_unit(uid,tp,lv,num) values(?,?,?,?)", uid, tp, lv, num)
		u = &protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}
	} else {
		Sql.Exec("update u_unit set num=num+? where uid=? and tp=? and lv=?;", num, uid, tp, lv)
		u.Num = proto.Int32(u.GetNum() + num)
	}
	a := []*protos.Unit{u}
	updates := &protos.Updates{}
	updates.Unit = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

func DelUnit(uid int64, tp int32, lv int32, num int32) {
	if num == 0 {
		return
	}
	u := ReadUnit(uid, tp, lv)
	if u == nil || u.GetNum() < num {
		Net.CallUidError(uid, 0, 1, "lessUnit")
		return
	} else {
		Sql.Exec("update u_unit set num=num-? where uid=? and tp=? and lv=?;", num, uid, tp, lv)
		u.Num = proto.Int32(u.GetNum() - num)
	}
	a := []*protos.Unit{u}
	updates := &protos.Updates{}
	updates.Unit = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

func ReadWounded(uid int64, tp int32, lv int32) *protos.Unit {
	rows, err := Sql.Query("select num from u_wounded where uid=? and tp=? and lv=?", uid)
	if err != nil {
		log.Println("City.OnUserInit error: ", err)
		return nil
	}
	for rows.Next() {
		var num int32
		rows.Scan(&num)
		rows.Close()
		return &protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}
	}
	return nil
}

func AddWounded(uid int64, tp int32, lv int32, num int32) {
	if num == 0 {
		return
	}
	u := ReadUnit(uid, tp, lv)
	if u == nil {
		Sql.Exec("replace into u_wounded(uid,tp,lv,num) values(?,?,?,?)", uid, tp, lv, num)
		u = &protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}
	} else {
		Sql.Exec("update u_unit set num=num+? where uid=? and tp=? and lv=?;", num, uid, tp, lv)
		u.Num = proto.Int32(u.GetNum() + num)
	}
	a := []*protos.Unit{u}
	updates := &protos.Updates{}
	updates.Wounded = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

func DelWounded(uid int64, tp int32, lv int32, num int32) {
	if num == 0 {
		return
	}
	u := ReadUnit(uid, tp, lv)
	if u == nil || u.GetNum() < num {
		Net.CallUidError(uid, 0, 1, "lessUnit")
		return
	} else {
		Sql.Exec("update u_wounded set num=num-? where uid=? and tp=? and lv=?;", num, uid, tp, lv)
		u.Num = proto.Int32(u.GetNum() - num)
	}
	a := []*protos.Unit{u}
	updates := &protos.Updates{}
	updates.Wounded = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

//----------------------------------------------------------
//event
func init() {
	Event.RegA("OnUserNew", func(uid int64) {
		Sql.Query(`insert into u_unit(uid,tp,lv,num) values(?,?,?,?)`,
			uid, 1, 1, 100)
		Sql.Query(`insert into u_wounded(uid,tp,lv,num) values(?,?,?,?)`,
			uid, 1, 1, 10)
	})
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		rows, err := Sql.Query("select * from u_unit where uid=?", uid)
		if err != nil {
			log.Println("City.OnUserInit error: ", err)
			return
		}
		a := []*protos.Unit{}
		for rows.Next() {
			var uid int64
			var tp, lv, num int32
			rows.Scan(&uid, &tp, &lv, &num)
			a = append(a, &protos.Unit{
				Tp:  proto.Int32(tp),
				Lv:  proto.Int32(lv),
				Num: proto.Int32(num),
			})
		}
		if len(a) == 0 {
			return
		}
		updates.Unit = a
	})
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		//伤兵
		rows, err := Sql.Query("select * from u_wounded where uid=?", uid)
		if err != nil {
			log.Println("City.OnUserInit error: ", err)
			return
		}
		a := []*protos.Unit{}
		for rows.Next() {
			var uid int64
			var tp, lv, num int32
			rows.Scan(&uid, &tp, &lv, &num)
			a = append(a, &protos.Unit{
				Tp:  proto.Int32(tp),
				Lv:  proto.Int32(lv),
				Num: proto.Int32(num),
			})
		}
		if len(a) == 0 {
			return
		}
		updates.Wounded = a
	})
}

//-----------------------------------------
//RPC
func init() {
	Net.RegRPC(Rpc.UnitTrain_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UnitTrain_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		tp, lv, num := ps.GetTp(), ps.GetLv(), ps.GetNum()
		fmt.Println(">>>UnitTrain_C", tp, lv, num)

		Item.DelRes(uid, 1, 1, "UnitTrain_C")
		if ps.GetImmed() { //秒
			Item.DelRes(uid, Item.RES_GEM, 1, "train")
			AddUnit(uid, tp, lv, num)
			return
		}

		stTime := (time.Now().UnixNano() / 1e6)
		edTime := stTime + 60000 //cfg.

		unit := []*protos.Unit{&protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}}
		units := &protos.UnitArray{Unit: unit}
		units_data, _ := proto.Marshal(units)

		_, sidj, _ := Sql.Exec(`insert into u_job(sid,uid,tp,stTime,edTime,bid,unit)
			values(null,?,?,?,?,?,?)`,
			uid, JOB_TRAIN, stTime, edTime, 0, string(units_data))

		job := []*protos.Job{&protos.Job{
			Sid:    proto.Int64(sidj),
			Tp:     proto.Int32(JOB_TRAIN),
			StTime: proto.Int64(stTime),
			EdTime: proto.Int64(edTime),
			Unit:   unit,
		}}
		updates := &protos.Updates{}
		updates.Job = job
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})
	Net.RegRPC(Rpc.UnitDisMiss_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UnitDisMiss_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		tp, lv, num := ps.GetTp(), ps.GetLv(), ps.GetNum()
		fmt.Println(">>>UnitDissMiss_C", tp, lv, num)
		DelUnit(uid, tp, lv, num)
	})
	Net.RegRPC(Rpc.UnitUp_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UnitUp_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		tp, lv, num := ps.GetTp(), ps.GetLv(), ps.GetNum()
		fmt.Println(">>>UnitUp_C", tp, lv, num)

		Item.DelRes(uid, 1, 1, "UnitTrain_C")
		DelUnit(uid, tp, lv, num)
		if ps.GetImmed() { //秒
			Item.DelRes(uid, Item.RES_GEM, 1, "trainFast")
			AddUnit(uid, tp, lv+1, num)
			return
		}

		stTime := (time.Now().UnixNano() / 1e6) //cfg. -
		edTime := stTime + 30000                //cfg. -

		unit := []*protos.Unit{&protos.Unit{
			Tp:  proto.Int32(tp),
			Lv:  proto.Int32(lv),
			Num: proto.Int32(num),
		}}
		units := &protos.UnitArray{Unit: unit}
		units_data, _ := proto.Marshal(units)

		_, sidj, _ := Sql.Exec(`insert into u_job(sid,uid,tp,stTime,edTime,bid,unit)
			values(null,?,?,?,?,?,?)`,
			uid, JOB_TRAINUP, stTime, edTime, 0, string(units_data))

		job := []*protos.Job{&protos.Job{
			Sid:    proto.Int64(sidj),
			Tp:     proto.Int32(JOB_TRAINUP),
			StTime: proto.Int64(stTime),
			EdTime: proto.Int64(edTime),
			Unit:   unit,
		}}
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: &protos.Updates{
				Job: job,
			},
		})
	})
	Net.RegRPC(Rpc.UnitHeal_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UnitHeal_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		unit := ps.GetUnit()
		fmt.Println(">>>UnitHeal_C", unit)
		if ps.GetImmed() { //秒
			Item.DelRes(uid, Item.RES_GEM, 1, "trainFast")
			for _, u := range unit {
				DelUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
				AddUnit(uid, u.GetTp(), u.GetLv()+1, u.GetNum())
			}
			return
		}
		//sum cost
		cost := []*protos.IdNum{}
		for _, u := range unit {
			//check num u.GetTp(),u.GetLv(),u.GetNum()
			log.Println("check", u.GetTp(), u.GetLv(), u.GetNum())
			//res,err:=
			//sum cost
			cost = append(cost, &protos.IdNum{
				Cid: proto.Int32(1),
				Num: proto.Int64(1),
			})
		}
		Item.DelRess(uid, cost, "heal")
		for _, u := range unit {
			DelUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
		}

	})

}
