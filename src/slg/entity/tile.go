package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//地板格
type Tile struct {
	X        int32 `xorm:"pk"`
	Y        int32 `xorm:"pk"`
	Area     int32 `xorm:"index(area)"`
	Tp       int32 `xorm:"index(type)"`
	Tp2      int32 `xorm:"index(type)"`
	Level    int32 `xorm:"index(type)"`
	Uid      int64 `xorm:"index(uid)"`
	Alliance int64 `xorm:"index(aid)"`
	Troops   []int64
	Version  int32 `xorm:"version"` //乐观锁
}

//转proto对象
func (this *Tile) ToProto() *protos.Tile {
	return &protos.Tile{
		X:   proto.Int32(this.X),
		Y:   proto.Int32(this.Y),
		Tp:  proto.Int32(this.Tp),
		Tp2: proto.Int32(this.Tp2),
	}
}

//加到更新
func (this *Tile) AppendTo(updates *protos.Updates) {
	list := updates.Tile
	if list == nil {
		list = []*protos.Tile{}
	}
	updates.Tile = append(list, this.ToProto())
}
