package World

import (
	"common/sql"
	"log"
	"math/rand"
	"slg/entity"
)

var Areas [AREA_ROWCOL][AREA_ROWCOL]*Entity.Area
var Tiles [WORLD_WIDTH][WORLD_WIDTH]*Entity.Tile
var Sights [SIGHT_ROWCOL][SIGHT_ROWCOL]*Entity.Sight

var ttt = make(chan func(), 512)

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
	var areaLoadOk = make(chan int, AREA_NUM)
	for r := 0; r < AREA_ROWCOL; r++ {
		for c := 0; c < AREA_ROWCOL; c++ {
			if Areas[r][c] == nil {
				go func(r, c int) {
					initArea(r, c)
					areaLoadOk <- 1
				}(r, c)
			} else {
				go func(r, c int) {
					loadArea(r, c)
					areaLoadOk <- 1
				}(r, c)
			}
		}
	}
	for i := 1; i <= AREA_NUM; i++ {
		<-areaLoadOk
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
	num := 10 //
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
