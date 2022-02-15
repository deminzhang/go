package Mail

import (
	"common/net"
	"common/sql"
	"common/utilX"
	"log"
	"protos"

	"slg/const"
	"slg/entity"
	"strings"
)

const (
	MAIL_TYPE_SYS    = 1 //系统
	MAIL_TYPE_USER   = 2 //玩家
	MAIL_TYPE_SCOUT  = 3 //侦查情报
	MAIL_TYPE_REPORT = 4 //战报
	MAIL_TYPE_ACTION = 5 //行动报告
)

func Read(uid int64, sid int64) *Entity.Mail {
	x := Sql.ORM()
	var mail Entity.Mail
	_, err := x.Where("Sid = ?", sid).Get(&mail)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &mail
}

func ReadFrom(uid int64, from int64) []Entity.Mail {
	x := Sql.ORM()

	mails := make([]Entity.Mail, 0)
	err := x.Where("Sid>=? and Uid = ?", from, uid).Find(&mails)
	if err != nil {
		log.Println(err)
		return nil
	}
	return mails
}

func SystemSend(uid int64, cid int32, titles []string, contexts []string, item []*protos.IdNum, res []*protos.IdNum) {
	now := utilX.MilliSecond()
	// timeOut := now + 7*24*86400000
	params := strings.Join(titles, ",")

	//TODO 用字串存库方便查库
	// its := &protos.ItemArray{Item: item}
	// items, _ := proto.Marshal(its)
	// ress := &protos.ResArray{Res: res}
	// resss, _ := proto.Marshal(ress)

	x := Sql.ORM()
	mail := Entity.Mail{
		Type:     MAIL_TYPE_SYS,
		FromUid:  0,
		FromName: "",
		Cid:      cid,
		Time:     now,
		Params:   params,
		// TimeOut:  timeOut,
		// ReportId: 0,
		// IntelId:  0,
		// Item:     item,
		// Res:      res,
		Read:  false,
		Take:  len(item)+len(res) == 0,
		Favor: false,
	}
	x.Insert(mail)

	updates := &protos.Updates{}
	mail.AppendTo(updates)
	Net.CallUid(uid, Const.Response_S, &protos.Response_S{
		ProtoId: 0,
		Updates: updates,
	})
}

func UserSend(uid int64, sender *Entity.User, title string, context string, item []*protos.IdNum, res []*protos.IdNum) {
	now := utilX.MilliSecond()
	// timeOut := now + 7*24*86400000
	x := Sql.ORM()
	mail := Entity.Mail{
		Type:     MAIL_TYPE_USER,
		FromUid:  sender.Uid,
		FromName: sender.Name,
		Cid:      0,
		Time:     now,
		Params:   title,
		// TimeOut:  timeOut,
		// ReportId: 0,
		// IntelId:  0,
		// Item:     item,
		// Res:      res,
		Read:  false,
		Take:  len(item)+len(res) == 0,
		Favor: false,
	}
	x.Insert(mail)

	updates := &protos.Updates{}
	mail.AppendTo(updates)
	Net.CallUid(uid, Const.Response_S, &protos.Response_S{
		ProtoId: 0,
		Updates: updates,
	})
}
