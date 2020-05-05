package Entity

import (
	"sync"
)

//世界地图分视区 内存数据不存库
type Sight struct {
	sync.RWMutex
	Sid    int32 //主键
	Row    int32
	Col    int32
	Eyes   map[int64]bool
	Tiles  []*Tile
	Nearby []*Sight
}
