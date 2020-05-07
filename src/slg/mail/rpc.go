package Mail

import (
	"common/net"
	"common/sql"
	"log"
	"protos"
	"slg/const"

	// "slg/item"

	// "strings"

	"github.com/golang/protobuf/proto"
)

//RPC
func init() {
	Net.RegRpcC(Const.MailGet_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.MailGet_C{}
		if err := proto.Unmarshal(data, &ps); err != nil {
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
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(int32(pid)),
			Updates: updates,
		})
	})
	{
		return
	}
	Net.RegRpcC(Const.MailDel_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.MailDel_C{}
		if err := proto.Unmarshal(data, &ps); err != nil {
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
		ss.CallOut(pid+1, &protos.Response_S{ProtoId: proto.Int32(int32(pid)),
			Removes: removes,
		})
	})
	Net.RegRpcC(Const.MailRead_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.MailRead_C{}
		if err := proto.Unmarshal(data, &ps); err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
	})
	Net.RegRpcC(Const.MailTake_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {

		ps := protos.MailTake_C{}
		if err := proto.Unmarshal(data, &ps); err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
	})
}
