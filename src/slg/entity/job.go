package Entity

import (
	"protos"
)

//工作队
type Job struct {
	Sid        int64 `xorm:"pk autoincr"`
	Uid        int64 `xorm:"index"`
	Tp         int32
	Create     int64 `xorm:"created"`
	St         int64
	Et         int64
	TotalTime  int64
	BuildingId int64
	TechId     int64
	UnitTp     int64
}

//转proto对象
func (this *Job) ToProto() *protos.Job {
	return &protos.Job{
		Sid:    this.Sid,
		Tp:     this.Tp,
		StTime: this.St,
		EdTime: this.Et,
	}
}

//转proto前端主键对象
func (this *Job) ToProtoPK() *protos.JobPK {
	return &protos.JobPK{
		Sid: this.Sid,
	}
}

//加到更新
func (this *Job) AppendTo(updates *protos.Updates) {
	list := updates.Job
	if list == nil {
		list = []*protos.Job{}
	}
	updates.Job = append(list, this.ToProto())
}

//加到删除
func (this *Job) AppendToPK(removes *protos.Removes) {
	list := removes.Job
	if list == nil {
		list = []*protos.JobPK{}
	}
	removes.Job = append(list, this.ToProtoPK())
}
