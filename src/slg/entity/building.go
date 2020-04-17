package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//建筑
type Building struct {
	Sid   int64 `xorm:"pk"`
	Uid   int64 `xorm:"index"`
	Cid   int32 `xorm:"index"`
	Level int32
	Pos   int32
}

//转proto对象
func (this *Building) ToProto() *protos.Building {
	return &protos.Building{
		Sid:   proto.Int64(this.Sid),
		Tp:    proto.Int32(this.Cid),
		Level: proto.Int32(this.Level),
		Pos:   proto.Int32(this.Pos),
	}
}

//转proto前端主键对象
func (this *Building) ToProtoPK() *protos.BuildingPK {
	return &protos.BuildingPK{
		Sid: proto.Int64(this.Sid),
	}
}

//加到更新
func (this *Building) AppendTo(updates *protos.Updates) {
	list := updates.Building
	if list == nil {
		list = []*protos.Building{}
	}
	updates.Building = append(list, this.ToProto())
}

//加到删除
func (this *Building) AppendToPK(removes *protos.Removes) {
	list := removes.Building
	if list == nil {
		list = []*protos.BuildingPK{}
	}
	removes.Building = append(list, this.ToProtoPK())
}
