package Item

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
		log.Println("Item.OnInitDB")
		x := Sql.ORM()
		x.Sync2(new(Entity.Item))
		x.Sync2(new(Entity.Res))
	})

	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("Item.OnUserNew", uid)
		var cid int32
		for cid = 1; cid < 10; cid++ {
			Sql.ORM().Insert(&Entity.Res{Uid: uid, Cid: cid, Num: 1})
		}

	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("Item.OnUserGetData", uid)
		x := Sql.ORM()

		items := make([]Entity.Item, 0)
		err := x.Where("Uid = ?", uid).Find(&items)
		if err != nil {
			for _, o := range items {
				o.AppendTo(updates)
			}
		}

		ress := make([]Entity.Res, 0)
		err = x.Where("Uid = ?", uid).Find(&ress)
		if err != nil {
			for _, o := range ress {
				o.AppendTo(updates)
			}
		}
	})
}
