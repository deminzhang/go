package Entity

import (
	"sync"
)

const (
	WORLD_MAIN_ID = 0    //主世界ID
	WORLD_WIDTH   = 1200 //世界宽
	//区域用于分区刷怪/建号
	AREA_ROWCOL   = 8                         //区域行列数
	AREA_NUM      = AREA_ROWCOL * AREA_ROWCOL //区域数
	AREA_WIDTH    = WORLD_WIDTH / AREA_ROWCOL //区域宽
	AREA_TILE_NUM = AREA_WIDTH * AREA_WIDTH   //区域格数
	//用于视野同步
	SIGHT_WIDTH  = 10                        //视块宽
	SIGHT_ROWCOL = WORLD_WIDTH / SIGHT_WIDTH //视野行列数
)

const (
	TILE_EMPTY   = 0 //空地
	TILE_MINE    = 1 //资源
	TILE_MONSTER = 2 //怪
)

type World struct {
	sync.RWMutex `xorm:"-"`
	Sid          int
	Inited       bool
}
