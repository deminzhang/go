package Mail

import (
	"common/event"
	"common/sql"
	"log"
	"protos"

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

		// SendSys(uid, 0, []string{"第一封"}, []string{"测试邮件"},
		// 	[]*protos.IdNum{
		// 		&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)},
		// 		&protos.IdNum{Cid: proto.Int32(2), Num: proto.Int64(2)},
		// 	},
		// 	[]*protos.IdNum{&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)}})
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("Mail.OnUserGetData", uid)
		// x := Sql.ORM()

		// items := make([]Entity.Item, 0)
		// err := x.Where("Uid = ?", uid).Find(&items)
		// if err != nil {
		// 	log.Println(err)
		// }
		// 	for _, o := range items {
		// 		o.AppendTo(updates)
		// 	}
	})
	// Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
	// 	var all, unread int64
	// 	rows, err := Sql.Query("select count(*) as count from u_mail where uid=?", uid)
	// 	if err != nil {
	// 		log.Println("Mail.OnUserInit error: ", err)
	// 		return
	// 	}
	// 	for rows.Next() {
	// 		rows.Scan(&all)
	// 	}
	// 	rows2, err := Sql.Query("select count(*) as count from u_mail where uid=? and `read`=0", uid)
	// 	if err != nil {
	// 		log.Println("Mail.OnUserInit error: ", err)
	// 		return
	// 	}
	// 	for rows2.Next() {
	// 		rows2.Scan(&unread)
	// 	}
	// 	a := []*protos.MailNum{&protos.MailNum{
	// 		Unread: proto.Int32(int32(all)),
	// 		All:    proto.Int32(int32(unread)),
	// 	}}
	// 	updates.MailNum = a
	// })
}
