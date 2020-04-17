package Mail

import (
	"common/net"
	"common/sql"
	"common/util"
	"log"
	"protos"

	// "slg/item"
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
			Content:  proto.String(string(context)),
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
			Content:  proto.String(string(context)),
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
	now := Util.MilliSecond()
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
		Content:  proto.String(context),
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
