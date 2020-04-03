package City

import (
	"common/event"
	"common/net"
	"common/sql"
	"fmt"
	"log"

	//	"net"
	"protos"
	"slg/item"
	"slg/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

func ReadBuilding(uid0 int64, sid int64) *protos.Building {
	rows, err := Sql.Query("select * from u_building where sid=?", sid)
	if err != nil {
		log.Println("ReadBuilding error: ", err)
		return nil
	}
	for rows.Next() {
		var sid, uid, resTime int64
		var tp, level, x, y, resOut, resVal int32
		rows.Scan(&sid, &uid, &tp, &level, &x, &y, &resTime, &resOut, &resVal)
		rows.Close()
		if uid0 != uid {
			return nil
		}
		return &protos.Building{
			Sid:   proto.Int64(sid),
			Tp:    proto.Int32(tp),
			Level: proto.Int32(level),
			X:     proto.Int32(x),
			Y:     proto.Int32(y),
			// ResTime: proto.Int64(resTime),
			// ResOut:  proto.Int32(resOut),
			// ResVal:  proto.Int32(resVal),
		}
	}
	return nil
}

//----------------------------------------------------------
//event
func init() {
	Event.RegA("OnUserNew", func(uid int64) {
		//for cfg if init level>0
		Sql.Query(`insert into u_building(sid,uid,tp,level,x,y,resTime,resOut,resVal)
			values(null,?,?,?,?,?,?,?,?)`,
			uid, 1, 1, 26, 26, 0, 0, 0)
	})
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		rows, err := Sql.Query("select * from u_building where uid=?", uid)
		if err != nil {
			log.Println("OnUserInit error: ", err)
			return
		}
		a := []*protos.Building{}
		for rows.Next() {
			var sid, uid, resTime int64
			var tp, level, x, y, resOut, resVal int32
			rows.Scan(&sid, &uid, &tp, &level, &x, &y, &resTime, &resOut, &resVal)
			a = append(a, &protos.Building{
				Sid:   proto.Int64(sid),
				Tp:    proto.Int32(tp),
				Level: proto.Int32(level),
				X:     proto.Int32(x),
				Y:     proto.Int32(y),
				// ResTime: proto.Int64(resTime),
				// ResOut:  proto.Int32(resOut),
				// ResVal:  proto.Int32(resVal),
			})
		}
		updates.Building = a
	})
}

//-----------------------------------------
//RPC
func init() {
	Net.RegRPC(Rpc.Build_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		fmt.Println(">>>Build_C", data)

		ps := protos.Build_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Build_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		tp, x, y := ps.GetTp(), ps.GetX(), ps.GetY()
		fmt.Println(">>>Build_C", tp, x, y)
		//TODO 建设条件...
		//
		// 最大个数
		// 建设耗费...
		Item.DelRes(uid, 1, 1, "build")
		updates := &protos.Updates{}
		var level int32 = 0
		var sid int64
		if ps.GetImmed() {
			Item.DelRes(uid, Item.RES_GEM, 1, "build")
			level = 1
			_, sid, _ = Sql.Exec(`insert into u_building(sid,uid,tp,level,x,y,resTime,resOut,resVal)
			values(null,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
				uid, tp, level, x, y, 0, 0, 0)

		} else {
			stTime := (time.Now().UnixNano() / 1e6)
			edTime := stTime + 60000 //cfg.

			_, sid, _ = Sql.Exec(`insert into u_building(sid,uid,tp,level,x,y,resTime,resOut,resVal)
			values(null,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
				uid, tp, level, x, y, 0, 0, 0)

			_, sidj, _ := Sql.Exec(`insert into u_job(sid,uid,tp,stTime,edTime,bid)
			values(null,?,?,?,?,?)`,
				uid, JOB_BUILD, stTime, edTime, sid)
			updates.Job = []*protos.Job{&protos.Job{
				Sid:    proto.Int64(sidj),
				Tp:     proto.Int32(JOB_BUILD),
				StTime: proto.Int64(stTime),
				EdTime: proto.Int64(edTime),
				Bid:    proto.Int64(sid),
			}}
		}
		updates.Building = []*protos.Building{&protos.Building{
			Sid:   &sid,
			Tp:    proto.Int32(tp),
			Level: proto.Int32(level),
			X:     proto.Int32(x),
			Y:     proto.Int32(y),
			// ResTime: proto.Int64(0),
			// ResOut:  proto.Int32(0),
			// ResVal:  proto.Int32(0),
		}}
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})
	Net.RegRPC(Rpc.BuildUp_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		fmt.Println(">>>BuildUp_C", data)

		ps := protos.BuildUp_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("BuildUp_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		sid := ps.GetSid()
		fmt.Println(">>>BuildUp_C", sid)

		b := ReadBuilding(uid, sid)
		if b == nil {
			ss.PError(protoId, 1, "nobuilding")
			return
		}
		//TODO
		level := b.GetLevel() //checkNext
		//j,_ := ReadJob()

		//TODO checkNextLevel Cost Condition

		Item.DelRes(uid, 1, 1, "buildUp")
		updates := &protos.Updates{}
		if ps.GetImmed() {
			Item.DelRes(uid, Item.RES_GEM, 1, "build")
			Sql.Exec(`update u_building set level=level+1 where sid=?`, sid)
			b.Level = proto.Int32(level + 1)
			updates.Building = []*protos.Building{b}
		} else {
			stTime := (time.Now().UnixNano() / 1e6)
			edTime := stTime + 60000 //cfg.
			_, sidj, _ := Sql.Exec(`insert into u_job(sid,uid,tp,stTime,edTime,bid)
			values(null,?,?,?,?,?)`,
				uid, JOB_BUILD, stTime, edTime, sid)

			updates.Job = []*protos.Job{&protos.Job{
				Sid:    proto.Int64(sidj),
				Tp:     proto.Int32(JOB_BUILD),
				StTime: proto.Int64(stTime),
				EdTime: proto.Int64(edTime),
				Bid:    proto.Int64(sid),
			}}
		}

		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})
	Net.RegRPC(Rpc.BuildMove_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.BuildMove_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("BuildMove_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		sid, x, y := ps.GetSid(), ps.GetX(), ps.GetY()
		fmt.Println(">>>BuildMove_C", sid, x, y)

		b := ReadBuilding(uid, sid)
		if b == nil {
			ss.PError(protoId, 1, "nobuilding")
			return
		}
		Sql.Exec(`update u_building set x=?, y=? where sid=?`, x, y, sid)
		b.X = proto.Int32(x)
		b.Y = proto.Int32(y)
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: &protos.Updates{Building: []*protos.Building{b}},
		})
	})

}
