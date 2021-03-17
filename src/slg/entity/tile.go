package Entity

import (
	"protos"
	"sync"
)

//地板格
type Tile struct {
	sync.RWMutex `xorm:"-"`
	X            int32   `xorm:"pk"`
	Y            int32   `xorm:"pk"`
	Area         int32   `xorm:"index(area)"` //所属区域
	Tp           int32   `xorm:"index(type)"` //类型
	Tp2          int32   `xorm:"index(type)"` //子类
	Level        int32   `xorm:"index(type)"` //级别
	Uid          int64   `xorm:"index(uid)"`  //所属玩家
	Alliance     int64   `xorm:"index(aid)"`  //所属联盟
	Troops       []int64 //停留部队
	Version      int32   `xorm:"version"` //乐观锁
}

//转proto对象
func (this *Tile) ToProto() *protos.Tile {
	return &protos.Tile{
		X:   this.X,
		Y:   this.Y,
		Tp:  this.Tp,
		Tp2: this.Tp2,
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

//加到更新
func AppendTo(a []Tile, updates *protos.Updates) {
	list := updates.Tile
	if list == nil {
		list = []*protos.Tile{}
	}
	for _, o := range a {
		updates.Tile = append(list, o.ToProto())
	}
}
