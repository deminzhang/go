package Entity

import (
	"sync"
)

//世界地图分视区
type Sight struct {
	sync.RWMutex
	Cid    int32
	Row    int32
	Col    int32
	Eyes   map[int64]bool
	Tiles  []*Tile
	Nearby []*Sight
}

//返回主键
func (this *Sight) GetPK() int32 {
	return this.Cid
}
