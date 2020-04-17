package World

import (
	"common/sql"
	"log"
	"math/rand"
	"slg/entity"
)

const (
	MAIN_WORLD  = 0    //主世界ID
	WORLD_WIDTH = 1200 //世界宽
	//区域用于分区刷怪/建号
	AREA_ROWCOL   = 8                         //区域行列数
	AREA_NUM      = AREA_ROWCOL * AREA_ROWCOL //区域数
	AREA_WIDTH    = WORLD_WIDTH / AREA_ROWCOL //区域宽
	AREA_TILE_NUM = AREA_WIDTH * AREA_WIDTH   //区域格数
	//用于视野同步
	SIGHT_WIDTH  = 10                        //视块宽
	SIGHT_ROWCOL = WORLD_WIDTH / SIGHT_WIDTH //视野行列数
)

var Areas [AREA_ROWCOL][AREA_ROWCOL]*Entity.Area
var Tiles [WORLD_WIDTH][WORLD_WIDTH]*Entity.Tile
var Sights [SIGHT_ROWCOL][SIGHT_ROWCOL]*Entity.Sight

var ttt = make(chan func(), 512)

var areaLoadCh = make(chan int, AREA_NUM)

func initWorld() {
	for y, line := range Tiles {
		for x, t := range line {
			if t == nil {
				Tiles[y][x] = &Entity.Tile{
					X:  int32(x),
					Y:  int32(y),
					Tp: 0,
				}
			}
		}
	}
	x := Sql.ORM()
	list := make([]Entity.Area, 0)
	err := x.Find(&list)
	if err != nil {
		log.Fatalln(err)
	}
	for _, o := range list {
		Areas[o.Row][o.Col] = &o
	}
	for r := 0; r < AREA_ROWCOL; r++ {
		for c := 0; c < AREA_ROWCOL; c++ {
			if Areas[r][c] == nil {
				go initArea(r, c)
			} else {
				go loadArea(r, c)
			}
		}
	}
	for i := 1; i <= AREA_NUM; i++ {
		<-areaLoadCh
	}
}

func initArea(r, c int) {
	log.Println("World.initArea", r, c)
	areaId := int32(r*AREA_ROWCOL + c)
	x := Sql.ORM()
	a := Entity.Area{
		Sid: areaId,
		Row: int32(r),
		Col: int32(c),
	}
	list := getEmptyTiles(r, c)
	num := 10
	for i := 1; i < num; i++ {
		tail := len(list) - i //1
		idx := rand.Intn(tail)
		sel := list[idx]
		list[idx] = list[tail]
		sel.Area = areaId
		sel.Tp = 1
		sel.Level = int32(rand.Intn(10)) + 1
		x.Insert(sel)
		list = list[:tail]
	}
	for i := 1; i < num; i++ {
		tail := len(list) - 1
		idx := rand.Intn(tail)
		sel := list[idx]
		list[idx] = list[tail]
		sel.Area = areaId
		sel.Tp = 2
		sel.Level = int32(rand.Intn(10)) + 1
		x.Insert(sel)
	}
	x.Insert(a)
	Areas[r][c] = &a
	areaLoadCh <- 1
}
func loadArea(r, c int) {
	log.Println("World.loadArea", r, c)
	areaId := int32(r*AREA_ROWCOL + c)
	x := Sql.ORM()
	list := make([]Entity.Tile, 0)
	err := x.Where("area = ?", areaId).Find(&list)
	if err != nil {
		log.Fatalln(err)
	}
	for _, o := range list {
		Tiles[o.Y][o.X] = &o
	}
	areaLoadCh <- 1
}

func getEmptyTiles(r, c int) []*Entity.Tile {
	sy := r * AREA_WIDTH
	sx := c * AREA_WIDTH
	list := make([]*Entity.Tile, 0)
	for y := sy; y < sy+AREA_WIDTH; y++ {
		for x := sx; x < sx+AREA_WIDTH; x++ {
			t := Tiles[y][x]
			if t != nil && t.Tp == 0 {
				list = append(list, t)
			}
		}
	}
	return list
}
