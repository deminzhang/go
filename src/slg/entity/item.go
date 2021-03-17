package Entity

import (
	"protos"
)

//道具实例
type Item struct {
	Sid  int64 `xorm:"pk autoincr"`
	Uid  int64 `xorm:"index"`
	Cid  int32 `xorm:"index"`
	Num  int64
	Time int64 `xorm:"updated"`
}

//转proto对象
func (this *Item) ToProto() *protos.Item {
	return &protos.Item{
		Sid: this.Sid,
		Cid: this.Cid,
		Num: this.Num,
	}
}

//转proto前端主键对象
func (this *Item) ToProtoPK() *protos.ItemPK {
	return &protos.ItemPK{
		Sid: this.Sid,
	}
}

//加到更新
func (this *Item) AppendTo(updates *protos.Updates) {
	list := updates.Item
	if list == nil {
		list = []*protos.Item{}
	}
	updates.Item = append(list, this.ToProto())
}

//加到删除
func (this *Item) AppendToPK(removes *protos.Removes) {
	list := removes.Item
	if list == nil {
		list = []*protos.ItemPK{}
	}
	removes.Item = append(list, this.ToProtoPK())
}
