package Mail

import (
	"common/net"
	"common/sql"
	"log"
	"protos"
	"slg/const"
	"slg/item"

	// "strings"

	"github.com/golang/protobuf/proto"
)

//RPC
func init() {
	Net.RegRPC(Const.MailGet_C, func(ss Net.Session, pid int32, data []byte, uid int64) {
		ps := protos.MailGet_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		fromSid := ps.GetFrom()
		updates := &protos.Updates{}
		mails := ReadFrom(uid, fromSid)
		for _, o := range mails {
			o.AppendTo(updates)
		}
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(pid),
			Updates: updates,
		})
	})
	{
		return
	}
	Net.RegRPC(Const.MailDel_C, func(ss Net.Session, pid int32, data []byte, uid int64) {
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
			return
		}
		removes := &protos.Removes{}
		removes.Mail = ms
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(pid),
			Removes: removes,
		})
	})
	Net.RegRPC(Const.MailRead_C, func(ss Net.Session, pid int32, data []byte, uid int64) {

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
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(pid),
			Props: updates,
		})
	})
	Net.RegRPC(Const.MailTake_C, func(ss Net.Session, pid int32, data []byte, uid int64) {

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
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(pid),
			Props: updates,
		})
	})
	Net.RegRPC(Const.MailFavor_C, func(ss Net.Session, pid int32, data []byte, uid int64) {

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
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(pid),
			Props: updates,
		})
	})
	Net.RegRPC(Const.ReadReport_C, func(ss Net.Session, pid int32, data []byte, uid int64) {

		ps := protos.ReadReport_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
	})
}
