package World

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
		log.Println("World.OnInitDB")
		x := Sql.ORM()
		x.Sync2(new(Entity.Area))
		x.Sync2(new(Entity.Tile))
	})
	Event.Reg(Const.OnLoadDB, func() {
		log.Println("World.OnLoadDB")
		initWorld(0)
	})
	Event.Reg(Const.OnServerStart, func() {
		log.Println("World.OnServerStart:startWorldTick")
		go worldTick()
	})
	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("World.OnUserNew", uid)

	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("World.OnUserGetData", uid)

	})
}
