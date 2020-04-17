package Unit

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
		x.Sync2(new(Entity.Unit))
	})
	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("Unit.OnUserNew", uid)
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("Unit.OnUserGetData", uid)
		x := Sql.ORM()
		list := make([]Entity.Unit, 0)
		err := x.Where("Uid = ?", uid).Find(&list)
		if err != nil {
			log.Println(err)
		}
		for _, o := range list {
			o.AppendTo(updates)
		}

	})
}
