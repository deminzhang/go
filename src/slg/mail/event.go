package Mail

import (
	"common/event"
	"common/sql"
	"log"

	"slg/const"
	"slg/entity"
	// "strings"
	// "github.com/golang/protobuf/proto"
)

//events
func init() {
	Event.Reg(Const.OnInitDB, func() {
		log.Println("Mail.OnInitDB")
		x := Sql.ORM()
		x.Sync2(new(Entity.Mail))
	})

	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("Mail.OnUserNew", uid)

		// SystemSend(uid, 0, []string{"第一封"}, []string{"测试邮件"},
		// 	[]*protos.IdNum{
		// 		&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)},
		// 		&protos.IdNum{Cid: proto.Int32(2), Num: proto.Int64(2)},
		// 	},
		// 	[]*protos.IdNum{&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)}})
	})
}
