package World

import (
	"common/sql"
	"log"

	// "math"
	"math/rand"
	"slg/entity"
	"sync"
)

type World struct {
	// sync.RWMutex
	Sid    int
	Inited bool

	areas_lock sync.RWMutex
	tile_lock  sync.RWMutex
	sight_lock sync.RWMutex
	areas      [AREA_ROWCOL][AREA_ROWCOL]*Entity.Area
	tiles      [WORLD_WIDTH][WORLD_WIDTH]*Entity.Tile
	sights     [SIGHT_ROWCOL][SIGHT_ROWCOL]*Entity.Sight
}

func (this *World) GetArea(row, col int) *Entity.Area {
	this.areas_lock.Lock()
	defer this.areas_lock.Unlock()
	return this.areas[row][col]
}

func (this *World) GetTile(x, y int) *Entity.Tile {
	this.tile_lock.Lock()
	defer this.tile_lock.Unlock()
	return this.tiles[x][y]
}

func (this *World) GetSight(row, col int) *Entity.Sight {
	this.sight_lock.Lock()
	defer this.sight_lock.Unlock()
	return this.sights[row][col]
}

func getAreaEmptyTiles(world *World, row, col int) []*Entity.Tile {
	sy := row * AREA_WIDTH
	sx := col * AREA_WIDTH
	list := make([]*Entity.Tile, 0)
	for y := sy; y < sy+AREA_WIDTH; y++ {
		for x := sx; x < sx+AREA_WIDTH; x++ {
			t := world.tiles[y][x]
			if t != nil && t.Tp == 0 {
				list = append(list, t)
			}
		}
	}
	return list
}

func getSightTiles(world *World, row, col int) []*Entity.Tile {
	sy := row * SIGHT_WIDTH
	sx := col * SIGHT_WIDTH
	list := make([]*Entity.Tile, 0)
	for y := sy; y < sy+SIGHT_WIDTH; y++ {
		for x := sx; x < sx+SIGHT_WIDTH; x++ {
			t := world.tiles[y][x]
			if t != nil && t.Tp != 0 {
				list = append(list, t)
			}
		}
	}
	log.Println(row, col, len(list))
	return list
}

func initArea(world *World, r, c int) {
	log.Println("World.initArea", r, c)
	areaId := int32(r*AREA_ROWCOL + c)
	x := Sql.ORM()
	a := Entity.Area{
		Sid: areaId,
		Row: int32(r),
		Col: int32(c),
	}
	list := getAreaEmptyTiles(world, r, c)
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
	world.areas[r][c] = &a
}

func loadArea(world *World, r, c int) {
	log.Println("World.loadArea", r, c)
	areaId := int32(r*AREA_ROWCOL + c)
	x := Sql.ORM()
	list := make([]Entity.Tile, 0)
	err := x.Where("area = ?", areaId).Find(&list)
	if err != nil {
		log.Fatalln(err)
	}
	for _, o := range list {
		world.tiles[o.Y][o.X] = &o
	}
}

//------------------------------------------------------------------------------
var AllWorlds = make(map[int]*World)

func initWorld(worldId int) {
	w := AllWorlds[worldId]
	if w != nil && w.Inited {
		log.Println("World had inited ", worldId)
		return
	}
	w = &World{
		Sid:    worldId,
		Inited: false,
	}

	for y, line := range w.tiles {
		for x, t := range line {
			if t == nil {
				w.tiles[y][x] = &Entity.Tile{
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
		w.areas[o.Row][o.Col] = &o
	}

	var areaLoadOk = make(chan int, AREA_NUM)
	for r := 0; r < AREA_ROWCOL; r++ {
		for c := 0; c < AREA_ROWCOL; c++ {
			if w.areas[r][c] == nil {
				go func(r, c int) {
					initArea(w, r, c)
					areaLoadOk <- 1
				}(r, c)
			} else {
				go func(r, c int) {
					loadArea(w, r, c)
					areaLoadOk <- 1
				}(r, c)
			}
		}
	}
	for i := 1; i <= AREA_NUM; i++ {
		<-areaLoadOk
	}
	w.Inited = true
	AllWorlds[worldId] = w
}

func moveEyes(worldId int, uid int64, x, y int32) []*Entity.Tile {
	// w := AllWorlds[worldId]
	// r := int(math.Floor(float64(y) / float64(SIGHT_WIDTH)))
	// c := int(math.Floor(float64(x) / float64(SIGHT_WIDTH)))
	r := int(y / SIGHT_WIDTH)
	c := int(x / SIGHT_WIDTH)
	w := AllWorlds[0]
	tiles := getSightTiles(w, r, c)
	return tiles
}

func GetEmptyTile(worldId, r, c int) *Entity.Tile {
	w := AllWorlds[worldId]
	sy := r * AREA_WIDTH
	sx := c * AREA_WIDTH
	//TODO 随机
	for y := sy; y < sy+AREA_WIDTH; y++ {
		for x := sx; x < sx+AREA_WIDTH; x++ {
			t := w.tiles[y][x]
			if t != nil && t.Tp == 0 {
				return t
			}
		}
	}
	return nil
}

//------------------------------------------------------------------------------
//chan or lock?
//独立操作world数据通道
var worldWorkQueue = make(chan func(), 1024)

func worldTick() {
	for {
		f := <-worldWorkQueue
		f()
	}
}

func PushWork(f func()) {
	worldWorkQueue <- f
}
