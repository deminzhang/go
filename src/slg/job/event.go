package Job

import (
	"common/event"
	"common/sql"
	"log"
	"protos"
	"slg/const"
	"slg/entity"
)

//event--------------------------------------
func init() {
	Event.Reg(Const.OnInitDB, func() {
		x := Sql.ORM()
		x.Sync2(new(Entity.Job))
	})
	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("Job.OnUserNew", uid)
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("Job.OnUserGetData", uid)

		x := Sql.ORM()
		list := make([]Entity.Job, 0)
		err := x.Where("uid = ?", uid).Find(&list)
		if err != nil {
			log.Println(err)
		}
		for _, o := range list {
			o.AppendTo(updates)
		}
	})

	// Event.Reg(Const.OnSecond, func(mills int64) {
	Event.Reg(Const.OnTick, func(mills int64) {
		// log.Println("Job.OnTick")
		x := Sql.ORM()
		list := make([]Entity.Job, 0)
		err := x.Where("et < ?", mills).Asc("et").Limit(100).Find(&list)
		if err != nil {
			log.Println(err)
		}
		for _, o := range list {
			Event.Call(Const.OnJobDone, o.Tp, &o)
			fn := _jobDoneCallBack[o.Tp]
			if fn != nil {
				fn(&o)
			} else {
				// log.Println("unknown job type :", o.Tp)
			}
		}
	})
	Event.Reg(Const.OnJobDone, func(tp int32, t *Entity.Job) {
		log.Println("OnJobDone", t)
		switch t {
		default:
			break
		}
		x := Sql.ORM()
		x.Delete(t)
	})
}
