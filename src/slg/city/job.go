package City

import (
	"common/event"
	"common/net"
	"common/sql"
	"fmt"
	"log"
	"protos"
	"slg/item"
	"slg/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	JOB_BUILD   = 1
	JOB_REMOVE  = 2
	JOB_TRAIN   = 3
	JOB_TRAINUP = 4
	JOB_HEAL    = 5
	JOB_RESERCH = 6
)

func onTick() {
	now := (time.Now().UnixNano() / 1e6)
	rows, _ := Sql.Query("select * from u_job where tp in(?,?) edTime>0 and edTime<=?", JOB_BUILD, JOB_REMOVE, now)
	for rows.Next() {
		var sid, uid, stTime, edTime, bid int64
		var tp int32
		var unit []byte
		rows.Scan(&sid, &uid, &tp, &stTime, &edTime, &bid, &unit)
		switch tp {
		case JOB_BUILD:

		case JOB_REMOVE:

		}
	}
}

func ReadJob(uid0 int64, sid int64) *protos.Job {
	rows, err := Sql.Query("select * from u_job where sid=?", sid)
	if err != nil {
		log.Println("City.OnUserInit error: ", err)
		return nil
	}
	for rows.Next() {
		var sid, uid, stTime, edTime, bid int64
		var tp int32
		var unit []byte
		rows.Scan(&sid, &uid, &tp, &stTime, &edTime, &bid, &unit)
		rows.Close()
		if uid0 != uid {
			return nil
		}
		ps := protos.UnitArray{}
		proto.Unmarshal(unit, &ps)
		return &protos.Job{
			Sid:    proto.Int64(sid),
			Tp:     proto.Int32(tp),
			StTime: proto.Int64(stTime),
			EdTime: proto.Int64(edTime),
			Bid:    proto.Int64(bid),
			Unit:   ps.GetUnit(),
		}
	}
	return nil
}

func SaveJob(uid int64, job *protos.Job) {
	units := &protos.UnitArray{Unit: job.GetUnit()}
	units_data, _ := proto.Marshal(units)
	Sql.Exec(`replace into u_job(sid,uid,tp,stTime,edTime,bid,unit)	values(?,?,?,?,?,?,?)`,
		job.GetSid(), uid, job.GetTp(), job.GetStTime(), job.GetEdTime(),
		job.GetBid(), string(units_data))
}

func DelJob(sid int64) {
	_, _, err := Sql.Exec("delete from u_job where sid=?", sid)
	if err != nil {
		log.Println("City.OnUserInit error: ", err)
	}
}

//----------------------------------------------------------
//event
func init() {
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		rows, err := Sql.Query("select * from u_job where uid=?", uid)
		if err != nil {
			log.Println("Job.OnUserInit error: ", err)
			return
		}
		job := []*protos.Job{}
		for rows.Next() {
			var sid, uid, stTime, edTime, bid int64
			var tp int32
			var unit []byte
			rows.Scan(&sid, &uid, &tp, &stTime, &edTime, &bid, &unit)

			ps := protos.UnitArray{}
			proto.Unmarshal(unit, &ps)
			job = append(job, &protos.Job{
				Sid:    proto.Int64(sid),
				Tp:     proto.Int32(tp),
				StTime: proto.Int64(stTime),
				EdTime: proto.Int64(edTime),
				Bid:    proto.Int64(bid),
				Unit:   ps.GetUnit(),
			})
		}
		updates.Job = job
	})
}

//-----------------------------------------
//RPC
func init() {
	Net.RegRPC(Rpc.JobDone_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.JobDone_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		sid := ps.GetSid()
		fmt.Println(">>>JobDone_C", sid)

		job := ReadJob(uid, sid)
		if job == nil {
			ss.PError(protoId, 1, "forbidden") //forbidden
			return
		}
		if job.GetEdTime() > (time.Now().UnixNano() / 1e6) {
			ss.PError(protoId, 1, "ing")
			return
		}
		DelJob(sid)
		switch job.GetTp() {
		case JOB_BUILD:
			break
		case JOB_REMOVE:
			break
		case JOB_TRAIN:
		case JOB_TRAINUP:
			u := job.GetUnit()[0]
			AddUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
			break
		case JOB_HEAL:
			for _, u := range job.GetUnit() {
				AddUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
			}
			break
		case JOB_RESERCH:
			break
		}

		removes := &protos.Removes{}
		removes.Job = []*protos.JobPK{&protos.JobPK{Sid: proto.Int64(sid)}}

		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			//Updates: updates,
			Removes: removes,
		})

	})
	Net.RegRPC(Rpc.JobCancel_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.JobCancel_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		sid := ps.GetSid()
		fmt.Println(">>>JobCancel_C", sid)

		job := ReadJob(uid, sid)
		if job == nil {
			ss.PError(protoId, 1, "forbidden") //forbidden
			return
		}
		if job.GetEdTime() < (time.Now().UnixNano() / 1e6) {
			ss.PError(protoId, 1, "done")
			return
		}
		DelJob(sid)
		//返还一半资源
		switch job.GetTp() {
		case JOB_BUILD:
		case JOB_REMOVE:
			break
		case JOB_TRAIN:
			for _, u := range job.GetUnit() {
				Item.AddRes(uid, 1, int64(u.GetNum()), "UnitUpCancel")
			}
		case JOB_TRAINUP:
			for _, u := range job.GetUnit() {
				Item.AddRes(uid, 1, int64(u.GetNum()), "UnitUpCancel")
			}
		case JOB_HEAL:
			for _, u := range job.GetUnit() {
				Item.AddRes(uid, 1, int64(u.GetNum()), "UnitUpCancel")
			}
		case JOB_RESERCH:
			break
		}
		//for _, u := range job.GetUnit() {
		//	Item.AddRes(uid, )
		//}

	})
	Net.RegRPC(Rpc.JobFast_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.JobFast_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		sid, item := ps.GetSid(), ps.GetItem()
		fmt.Println(">>>JobFast_C", sid, item)

		job := ReadJob(uid, sid)
		if job == nil {
			ss.PError(protoId, 1, "forbidden") //forbidden
			return
		}

		if ps.GetImmed() { //秒
			Item.DelRes(uid, Item.RES_GEM, 1, "trainFast")
			u := job.GetUnit()[0]
			AddUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
			DelJob(sid)
			return
		}
		// Item.DelRes(uid, 1, 1, "JobFast_C")
		// job.EdTime = proto.Int64(job.GetEdTime() - 10000)
		// now := (time.Now().UnixNano() / 1e6)
		// if job.GetEdTime() < now {
		// 	switch job.GetTp() {
		// 	case JOB_BUILD:
		// 	case JOB_REMOVE:
		// 		break
		// 	case JOB_TRAIN:
		// 	case JOB_TRAINUP:
		// 	case JOB_HEAL:
		// 	case JOB_RESERCH:
		// 		break
		// 	}
		// 	// u := job.GetUnit()[0]
		// 	// AddUnit(uid, u.GetTp(), u.GetLv(), u.GetNum())
		// 	DelJob(sid)
		// 	return
		// }
		// SaveJob(uid, job)
		// ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
		// 	Updates: &protos.Updates{
		// 		Job: []*protos.Job{job},
		// 	},
		// })
	})

}
