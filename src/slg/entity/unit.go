package Entity

import (
	"protos"
)

//资源计数
type Unit struct {
	Uid int64 `xorm:"pk"`
	Cid int32 `xorm:"pk"`
	Num int64
}

//转proto对象
func (this *Unit) ToProto() *protos.Unit {
	return &protos.Unit{
		Cid: this.Cid,
		Num: this.Num,
	}
}

//转proto前端主键对象
func (this *Unit) ToProtoPK() *protos.UnitPK {
	return &protos.UnitPK{
		Cid: this.Cid,
	}
}

//加到更新
func (this *Unit) AppendTo(updates *protos.Updates) {
	list := updates.Unit
	if list == nil {
		list = []*protos.Unit{}
	}
	updates.Unit = append(list, this.ToProto())
}

//加到删除
func (this *Unit) AppendToPK(removes *protos.Removes) {
	list := removes.Unit
	if list == nil {
		list = []*protos.UnitPK{}
	}
	removes.Unit = append(list, this.ToProtoPK())
}
