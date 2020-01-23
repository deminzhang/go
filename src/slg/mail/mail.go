package Mail

import (
	"common/event"
	"common/net"
	"common/sql"
	"common/util"
	"log"
	"protos"
	"slg/item"
	"slg/rpc"
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	Sys    = 1
	User   = 2
	SCOUT  = 3
	Report = 4
	Action = 5
)

//1系统2玩家3侦查情报4战报5行动报告

func Read(uid0 int64, sid int64) *protos.Mail {
	rows, err := Sql.Query(`select * from u_mail where sid=?`, sid)
	if err != nil {
		log.Println("Mail.OnUserInit error: ", err)
		return nil
	}
	for rows.Next() {
		var sid, uid, fromUid, time, timeOut, reportId, intelId, pubMail int64
		var tp, cid int32
		var fromName, title, context, item, res []byte
		var read, take, favor int8
		rows.Scan(&sid, &uid, &tp, &fromUid, &fromName, &cid, &title, &context, &time,
			&timeOut, &reportId, &intelId, &item, &res, &read, &take, &favor, &pubMail)
		rows.Close()
		if uid0 != uid {
			return nil
		}

		ps := protos.ItemArray{}
		proto.Unmarshal(item, &ps)
		ps2 := protos.ResArray{}
		proto.Unmarshal(res, &ps2)

		return &protos.Mail{
			Sid:      proto.Int64(sid),
			Tp:       proto.Int32(tp),
			FromUid:  proto.Int64(fromUid),
			FromName: proto.String(string(fromName)),
			Cid:      proto.Int32(cid),
			Title:    proto.String(string(title)),
			Context:  proto.String(string(context)),
			Time:     proto.Int64(time),
			TimeOut:  proto.Int64(timeOut),
			ReportId: proto.Int64(reportId),
			IntelId:  proto.Int64(intelId),
			Item:     ps.GetItem(),
			Res:      ps2.GetRes(),
			Read:     proto.Bool(read == 1),
			Take:     proto.Bool(take == 1),
			Favor:    proto.Bool(favor == 1),
		}
	}
	return nil
}

func Reads(uid int64, from int64) []*protos.Mail {
	rows, err := Sql.Query(`select * from u_mail where uid=? and sid>?`, uid, from)
	if err != nil {
		log.Println("Mail.OnUserInit error: ", err)
		return nil
	}
	columns, _ := rows.Columns()
	log.Println("Mail.columns: ", columns)
	a := []*protos.Mail{}
	for rows.Next() {
		var sid, uid, fromUid, time, timeOut, reportId, intelId, pubMail int64
		var tp, cid int32
		var fromName, title, context, item, res []byte
		var read, take, favor int8
		rows.Scan(&sid, &uid, &tp, &fromUid, &fromName, &cid, &title, &context, &time,
			&timeOut, &reportId, &intelId, &item, &res, &read, &take, &favor, &pubMail)

		ps := protos.ItemArray{}
		proto.Unmarshal(item, &ps)
		ps2 := protos.ResArray{}
		proto.Unmarshal(res, &ps2)

		a = append(a, &protos.Mail{
			Sid:      proto.Int64(sid),
			Tp:       proto.Int32(tp),
			FromUid:  proto.Int64(fromUid),
			FromName: proto.String(string(fromName)),
			Cid:      proto.Int32(cid),
			Title:    proto.String(string(title)),
			Context:  proto.String(string(context)),
			Time:     proto.Int64(time),
			TimeOut:  proto.Int64(timeOut),
			ReportId: proto.Int64(reportId),
			IntelId:  proto.Int64(intelId),
			Item:     ps.GetItem(),
			Res:      ps2.GetRes(),
			Read:     proto.Bool(read == 1),
			Take:     proto.Bool(take == 1),
			Favor:    proto.Bool(favor == 1),
		})
	}
	return a
}

func SendSys(uid int64, cid int32, titles []string, contexts []string, item []*protos.IdNum, res []*protos.IdNum) {
	now := Util.MSec()
	timeOut := now + 7*24*86400000
	title := strings.Join(titles, ",")
	context := strings.Join(contexts, ",")

	//TODO 用字串存库方便查库
	its := &protos.ItemArray{Item: item}
	items, _ := proto.Marshal(its)
	ress := &protos.ResArray{Res: res}
	resss, _ := proto.Marshal(ress)

	_, sid, _ := Sql.Exec(`insert into u_mail(sid,uid,tp,fromUid,cid,title,context,time,timeOut,item,res)
	values(null,?,?,?,?,?,?,?,?,?,?)`,
		uid, 1, 0, cid, title, context, now, timeOut, items, resss)

	m := &protos.Mail{
		Sid:      proto.Int64(sid),
		Tp:       proto.Int32(1),
		FromUid:  proto.Int64(0),
		FromName: proto.String(""),
		Cid:      proto.Int32(cid),
		Title:    proto.String(title),
		Context:  proto.String(context),
		Time:     proto.Int64(now),
		TimeOut:  proto.Int64(timeOut),
		ReportId: proto.Int64(0),
		IntelId:  proto.Int64(0),
		Item:     item,
		Res:      res,
		Read:     proto.Bool(false),
		Take:     proto.Bool(len(item)+len(res) == 0),
		Favor:    proto.Bool(false),
	}

	updates := &protos.Updates{}
	updates.Mail = []*protos.Mail{m}
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

//events
func init() {
	Event.RegA("OnUserNew", func(uid int64) {
		SendSys(uid, 0, []string{"第一封"}, []string{"测试邮件"},
			[]*protos.IdNum{
				&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)},
				&protos.IdNum{Cid: proto.Int32(2), Num: proto.Int64(2)},
			},
			[]*protos.IdNum{&protos.IdNum{Cid: proto.Int32(1), Num: proto.Int64(1)}})
	})
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		var all, unread int64
		rows, err := Sql.Query("select count(*) as count from u_mail where uid=?", uid)
		if err != nil {
			log.Println("Mail.OnUserInit error: ", err)
			return
		}
		for rows.Next() {
			rows.Scan(&all)
		}
		rows2, err := Sql.Query("select count(*) as count from u_mail where uid=? and `read`=0", uid)
		if err != nil {
			log.Println("Mail.OnUserInit error: ", err)
			return
		}
		for rows2.Next() {
			rows2.Scan(&unread)
		}
		a := []*protos.MailNum{&protos.MailNum{
			Unread: proto.Int32(int32(all)),
			All:    proto.Int32(int32(unread)),
		}}
		updates.MailNum = a
	})
}

//RPC
func init() {
	Net.RegRPC(Rpc.MailGet_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {

		ps := protos.MailGet_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		from := ps.GetFrom()
		updates := &protos.Updates{}
		updates.Mail = Reads(uid, from)
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})
	Net.RegRPC(Rpc.MailDel_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.MailDel_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		force := ps.GetForce()
		sids := ps.GetSids()

		ms := []*protos.MailPK{}
		if force {
			for _, sid := range sids {
				Sql.Exec("delete from u_mail where sid=?", sid)
				ms = append(ms, &protos.MailPK{Sid: proto.Int64(sid)})
			}
		} else {
			for _, sid := range sids {
				a, _, _ := Sql.Exec("delete from u_mail where sid=? and take=1", sid)
				if a > 0 {
					ms = append(ms, &protos.MailPK{Sid: proto.Int64(sid)})
				}
			}
		}
		if len(ms) == 0 {
			ss.PError(protoId, 1, "noMailDel")
			return
		}
		removes := &protos.Removes{}
		removes.Mail = ms
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Removes: removes,
		})
	})
	Net.RegRPC(Rpc.MailRead_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {

		ps := protos.MailRead_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		sids := ps.GetSids()
		a := []*protos.Mail{}
		for _, sid := range sids {
			af, _, _ := Sql.Exec("update u_mail set `read`=1 where sid=? and uid=?", sid, uid)
			if af == 0 {
				continue
			}
			m := &protos.Mail{
				Sid:  proto.Int64(sid),
				Read: proto.Bool(true),
			}
			a = append(a, m)
		}
		if len(a) == 0 {
			return
		}
		updates := &protos.Updates{}
		updates.Mail = a
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})
	Net.RegRPC(Rpc.MailTake_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {

		ps := protos.MailTake_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		sids := ps.GetSids()
		a := []*protos.Mail{}
		for _, sid := range sids {
			af, _, _ := Sql.Exec("update u_mail set take=1 where sid=? and uid=?", sid, uid)
			if af == 0 {
				continue
			}
			m := &protos.Mail{
				Sid:  proto.Int64(sid),
				Take: proto.Bool(true),
			}
			Item.Adds(uid, m.GetItem(), "mail")
			Item.AddRess(uid, m.GetRes(), "mail")
			a = append(a, m)
		}
		if len(a) == 0 {
			return
		}
		updates := &protos.Updates{}
		updates.Mail = a
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})
	Net.RegRPC(Rpc.MailFavor_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {

		ps := protos.MailFavor_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		sid := ps.GetSid()
		af, _, _ := Sql.Exec("update u_mail set favor=1 where sid=? and uid=?", sid, uid)
		if af == 0 {
			return
		}
		m := &protos.Mail{
			Sid:   proto.Int64(sid),
			Favor: proto.Bool(true),
		}
		updates := &protos.Updates{}
		updates.Mail = []*protos.Mail{m}
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})
	Net.RegRPC(Rpc.ReadReport_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {

		ps := protos.ReadReport_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
	})
}
