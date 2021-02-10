package Entity

import (
	"sync"
)

//世界地图分区
type Area struct {
	sync.RWMutex `xorm:"-"`
	Sid          int32   `xorm:"pk"`
	Col          int32   `xorm:"index(pos)"`
	Row          int32   `xorm:"index(pos)"`
	Inited       bool    `xorm:"-"`
	Tiles        []*Tile `xorm:"-"` //包含格子
	Version      int32   `xorm:"version"`
}
