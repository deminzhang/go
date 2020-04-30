package Troop

import (
	"common/event"
	"common/sql"
	"log"
	"protos"
	"slg/const"
	"slg/entity"
)

//event--------------------------------------
func init() {
	Event.Reg(Const.OnInitDB, func() {
		x := Sql.ORM()
		x.Sync2(new(Entity.Troop))
	})
	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("Troop.OnUserNew", uid)
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("Troop.OnUserGetData", uid)

		x := Sql.ORM()
		list := make([]Entity.Troop, 0)
		err := x.Where("uid = ?", uid).Find(&list)
		if err != nil {
			log.Println(err)
		}
		for _, o := range list {
			o.AppendTo(updates)
		}

	})

	// Event.Reg(Const.OnSecond, func(mills int64) {
	Event.Reg(Const.OnTick, func(mills int64) {
		// log.Println("Troop.OnTick")
		x := Sql.ORM()
		list := make([]Entity.Troop, 0)
		err := x.Where("et < ?", mills).Asc("et").Limit(100).Find(&list)
		if err != nil {
			log.Println(err)
		}
		for _, o := range list {
			Event.Call(Const.OnTroopStatDone, o.Tp, &o)
		}
	})
	Event.Reg(Const.OnTroopStatDone, func(tp int32, t *Entity.Troop) {
		log.Println("OnTroopStatDone", t)
		switch t.Tp {
		default:
			break
		}
		x := Sql.ORM()
		x.Delete(t)
	})

}
