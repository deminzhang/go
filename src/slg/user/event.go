package User

import (
	"common/event"
	"common/sql"
	"slg/const"
	"slg/entity"
	// "protos"
)

//event--------------------------------------
func init() {
	Event.Reg(Const.OnDBConnect, func() {
		x := Sql.ORM()
		x.Sync2(new(Entity.User))
	})
	// Event.Reg(Const.OnUserNew, func(uid int64) {

	// })
	// Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {

	// })
}
