package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//资源计数
type Res struct {
	Uid int64 `xorm:"pk"`
	Cid int32 `xorm:"pk"`
	Num int64
}

//转proto对象
func (this *Res) ToProto() *protos.Res {
	return &protos.Res{
		Cid: proto.Int32(this.Cid),
		Num: proto.Int64(this.Num),
	}
}

//转proto前端主键对象
func (this *Res) ToProtoPK() *protos.ResPK {
	return &protos.ResPK{
		Cid: proto.Int32(this.Cid),
	}
}

//加到更新
func (this *Res) AppendTo(updates *protos.Updates) {
	list := updates.Res
	if list == nil {
		list = []*protos.Res{}
	}
	updates.Res = append(list, this.ToProto())
}

//加到删除
func (this *Res) AppendToPK(removes *protos.Removes) {
	list := removes.Res
	if list == nil {
		list = []*protos.ResPK{}
	}
	removes.Res = append(list, this.ToProtoPK())
}
