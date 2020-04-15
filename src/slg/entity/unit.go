package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//资源计数
type Unit struct {
	Uid int64 `xorm:"pk"`
	Cid int32 `xorm:"pk"`
	Num int32
}

//返回主键
func (this *Unit) GetPK() (int64, int32) {
	return this.Uid, this.Cid
}

//转proto对象
func (this *Unit) ToProto() *protos.Unit {
	return &protos.Unit{
		Cid: proto.Int32(this.Cid),
		Num: proto.Int32(this.Num),
	}
}

//转proto前端主键对象
func (this *Unit) ToProtoPK() *protos.UnitPK {
	return &protos.UnitPK{
		Cid: proto.Int32(this.Cid),
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
