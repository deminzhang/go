package World

import (
	"common/event"
	"common/net"
	"common/sql"
	"fmt"
	"math"
	"math/rand"
	"protos"
	"slg/item"
	"slg/rpc"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

//!方便与lua对,数组从1开始
const (
	TILEX_NUM_X = 1200 //地图长
	TILEX_NUM_Y = 1200 //地图宽
	//区域用于分区刷怪/建号
	AREA_WIDTH    = 150                      //区域宽
	AREA_ROWN     = TILEX_NUM_Y / AREA_WIDTH //区域Y个数
	AREA_COLN     = TILEX_NUM_X / AREA_WIDTH //区域X个数
	AREA_TILE_NUM = AREA_WIDTH * AREA_WIDTH  //区域坐标个数
	//用于视野同步
	SIGHT_WIDTH = 10 //视块宽
	SIGHT_ROWN  = TILEX_NUM_Y / SIGHT_WIDTH
	SIGHT_COLN  = TILEX_NUM_X / SIGHT_WIDTH

	TILE_NONE = 0   //空地
	TILE_RES1 = 1   //资源1
	TILE_RES2 = 2   //资源2
	TILE_RES3 = 3   //资源3
	TILE_RES4 = 4   //资源4
	TILE_CITY = 999 //玩家
)

type Tile struct {
	Y   int
	X   int
	Tp  int
	Val int64
}

type Area struct {
	id     int //用于生成玩家选区
	col    int
	row    int
	inited bool
	citys  []int
	tiles  []*Tile
}

type Eye struct {
	uid    int64
	server int
	row    int
	col    int
}

type Sight struct {
	sync.RWMutex
	id     int
	row    int
	col    int
	inited bool
	eyes   map[int64]bool
	tiles  []*Tile
	nearby []*Sight
}

func (m *Sight) Set(k int64, v bool) {
	m.Lock()
	if v {
		m.eyes[k] = v
	} else {
		delete(m.eyes, k)
	}
	m.Unlock()
}
func (m *Sight) Get(k int64) bool {
	m.Lock()
	defer m.Unlock()
	return m.eyes[k]
}

var Tiles [TILEX_NUM_X + 1][TILEX_NUM_Y + 1]Tile
var Areas [AREA_ROWN + 1][AREA_COLN + 1]Area
var Sights [SIGHT_ROWN][SIGHT_COLN]*Sight

type EyesMap struct {
	sync.RWMutex
	list map[int64]*Eye
}

func (m *EyesMap) Set(k int64, v *Eye) {
	m.Lock()
	m.list[k] = v
	m.Unlock()
}
func (m *EyesMap) Get(k int64) *Eye {
	m.Lock()
	defer m.Unlock()
	return m.list[k]
}

var Eyes = &EyesMap{list: make(map[int64]*Eye)}

//var ServerSight map[int] Sight

func init() {
	rand.Seed(time.Now().UnixNano())

	// initTiles()
	// initAreas()
}

func (a *Area) init() {
	sy := (a.row-1)*AREA_WIDTH + 1
	sx := (a.col-1)*AREA_WIDTH + 1

	tx, err := Sql.Begin()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	rows, err := tx.Query("select * from w_area where row=? and col=?", a.row, a.col)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	for rows.Next() {
		rows.Close()

		rows, err = tx.Query("select x,y,tp,val from w_tile where x>=? and x<? and y>=? and y<?",
			sx, sx+AREA_WIDTH, sy, sy+AREA_WIDTH)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		n := 0
		for rows.Next() {
			var x, y, tp int
			var val int64
			rows.Scan(&x, &y, &tp, &val)
			tile := Tiles[y][x]
			tile.Tp = tp
			tile.Val = val
			Tiles[y][x] = tile
			n++
		}
		a.inited = true
		//fmt.Println(">>>readArea:", a.row, a.col, n)
		return
	}

	//fmt.Println(">>>initArea:row,col", a.row,a.col)
	//res20% mon10% block10% TODO
	//Config.global.res_propotion
	//Config.global.monster_propotion
	resNum := AREA_TILE_NUM * 20 / 100
	resIdx := rand.Perm(resNum)
	var list []*Tile
	block := false
	for y := sy; y < sy+AREA_WIDTH; y++ {
		for x := sx; x < sx+AREA_WIDTH; x++ {
			tile := Tiles[y][x]
			if tile.Tp == 0 && !block {
				list = append(list, &tile)
			}
		}
	}
	for _, i := range resIdx {
		tile := *list[i]
		tp := rand.Intn(4) + 1
		tile.Tp = tp
		if tp > 0 {
			tile.Val = int64(rand.Intn(10) + 1)
		}
		//TODO 暂不存,每次重启重生成
		tx.Exec("replace into w_tile(x,y,tp,val) values(?,?,?,?)", tile.X, tile.Y, tp, tile.Val)
	}
	monNum := AREA_TILE_NUM * 10 / 100
	//monIdx := rand.Perm(monNum)

	//TODO 暂不存,每次重启重生成
	tx.Exec("replace into w_area(row,col) values(?,?)", a.row, a.col)
	tx.Commit()
	fmt.Println(">>>initArea:row,col,res,mon", a.row, a.col, resNum, monNum)
}

func initTiles() {
	//TODO readBlock
	for y := 1; y <= TILEX_NUM_Y; y++ {
		for x := 1; x <= TILEX_NUM_X; x++ {
			Tiles[y][x].X = x
			Tiles[y][x].Y = y
		}
	}
}

func initAreas() {
	var idx int
	ch := make(chan int, AREA_ROWN*AREA_COLN)
	for r := 1; r <= AREA_ROWN; r++ {
		for c := 1; c <= AREA_COLN; c++ {
			idx++
			a := Areas[r][c]
			a.id = idx
			a.col = c
			a.row = r
			//go
			ch <- 1
			a.init()
		}
	}
	for i := 1; i <= AREA_ROWN*AREA_COLN; i++ {
		<-ch
	}
}

//---------------------------------------------------------
//Sight
func (s *Sight) init(row, col int) {
	Sights[row][col] = s
	s.id = row*10000 + col
	s.row = row
	s.col = col
	s.eyes = make(map[int64]bool)
	for y := (row-1)*SIGHT_WIDTH + 1; y <= row*SIGHT_WIDTH; y++ {
		for x := (col-1)*SIGHT_WIDTH + 1; x <= col*SIGHT_WIDTH; x++ {
			s.tiles = append(s.tiles, &Tiles[y][x])
		}
	}
	//fmt.Println(">>>Sight.init", a.tiles)
}
func (s *Sight) addEyes(uid int64) {
	Eyes.Set(uid, &Eye{uid: uid, row: s.row, col: s.col})
	s.eyes[uid] = true
}
func (s *Sight) delEyes(uid int64) {
	delete(s.eyes, uid)
	Eyes.Set(uid, nil)
}

//相邻九宫
func (s *Sight) allSights() []*Sight {
	t := []*Sight{}
	//return append(t, s) //暂用单宫
	//TODO 待优化为用nearby
	for r := s.row - 1; r <= s.row+1; r++ {
		if r > 0 && r <= SIGHT_ROWN {
			for c := s.col - 1; c <= s.col+1; c++ {
				if c > 0 && c <= SIGHT_COLN {
					aa := getSight(r, c)
					t = append(t, aa)
				}
			}
		}
	}
	return t
}
func (s *Sight) ForEachTile(foo func(int, int, int, int64)) {
	for _, tile := range s.tiles {
		t := *tile
		foo(t.X, t.Y, t.Tp, t.Val)
	}
}
func (s *Sight) CallRound(pid int, msg proto.Message) {
	for _, ss := range s.allSights() {
		for uid, _ := range ss.eyes {
			Net.CallUid(uid, int32(pid), msg)
		}
	}
}

//指定视区
func getSight(row, col int) *Sight {
	s := Sights[row][col]
	if s == nil {
		s = &Sight{}
		s.init(row, col)
	}
	return s
}

func getSightByXY(x, y int32) *Sight {
	row := int(math.Ceil(float64(y) / float64(SIGHT_WIDTH)))
	col := int(math.Ceil(float64(x) / float64(SIGHT_WIDTH)))
	s := getSight(row, col)
	return s
}

//重置视区 TODO跨服
func ResetSight(ss Net.Session, server int32, x int32, y int32) []*protos.Tile {
	uid := ss.GetUid()
	oldRow := 0 //ss.SightR
	oldCol := 0 //ss.SightC
	newRow := int(math.Ceil(float64(y) / float64(SIGHT_WIDTH)))
	newCol := int(math.Ceil(float64(x) / float64(SIGHT_WIDTH)))
	//if serverOld TODO 判定是否本服
	fmt.Println(">>>ResetSight", uid, oldRow, oldCol, newRow, newCol)
	if oldRow == newRow && oldCol == newCol { //视野区未变
		return nil
	}
	news := getSight(newRow, newCol)
	news.addEyes(uid)

	olds := make(map[int]bool)
	if oldRow != 0 && oldCol != 0 { //旧区
		old := getSight(oldRow, oldCol)
		for _, s := range old.allSights() {
			olds[s.id] = true
		}
		old.delEyes(uid)
	}
	tiles := []*protos.Tile{}
	for _, s := range news.allSights() {
		if !olds[s.id] {
			s.ForEachTile(func(x int, y int, tp int, val int64) {
				if tp == 0 {
					return
				}
				tiles = append(tiles, &protos.Tile{
					X:   proto.Int32(int32(x)),
					Y:   proto.Int32(int32(y)),
					Tp:  proto.Int32(int32(tp)),
					Val: proto.Int64(val),
				})
			})
		}
	}
	//ss.SightR, ss.SightC = newRow, newCol
	return tiles
}

//检查可迁入
func CheckCityLand(x, y int32) bool {
	//0 0 !5
	//0 0 !5
	//!5!5!5
	if Tiles[y][x].Tp > TILE_NONE || Tiles[y-1][x-1].Tp > TILE_NONE ||
		Tiles[y][x-1].Tp > TILE_NONE || Tiles[y-1][x].Tp > TILE_NONE {
		return false
	}
	if Tiles[y+1][x-1].Tp == TILE_CITY || Tiles[y+1][x].Tp == TILE_CITY ||
		Tiles[y+1][x+1].Tp == TILE_CITY || Tiles[y-1][x+1].Tp == TILE_CITY ||
		Tiles[y][x+1].Tp == TILE_CITY {
		return false
	}
	return true
}

//迁入
func MoveCity(uid int64, x, y int32) {
	var oldx, oldy int32
	rows, err := Sql.Query("select cityX,cityY from u_user where uid=?", uid)
	for rows.Next() {
		err = rows.Scan(&oldx, &oldy)
		if err != nil {
			panic(err)
		}
	}
	if oldx == x && oldy == y {
		return
	}
	Tiles[y][x].Tp = TILE_CITY
	Tiles[y][x].Val = uid
	tx, _ := Sql.Begin()
	tx.Exec("delete from w_tile where x=? and y=? and tp=? and val=?", oldx, oldy, TILE_CITY, uid)
	tx.Exec("replace into w_tile(x,y,tp,val) values(?,?,?,?)", x, y, TILE_CITY, uid)
	tx.Exec("update u_user set cityX=?,cityY=? where uid=?", x, y, uid)
	tx.Commit()
	//for old update tile
	olds := getSightByXY(oldx, oldy)
	tilerm := []*protos.TilePK{}
	tilerm = append(tilerm, &protos.TilePK{
		X: proto.Int32(int32(oldx)),
		Y: proto.Int32(int32(oldy)),
	})
	news := getSightByXY(x, y)
	//for new update tile
	tiles := []*protos.Tile{}
	tiles = append(tiles, &protos.Tile{
		X:   proto.Int32(int32(x)),
		Y:   proto.Int32(int32(y)),
		Tp:  proto.Int32(int32(TILE_CITY)),
		Val: proto.Int64(uid),
	})
	if news == olds {
		news.CallRound(12, &protos.Response_S{ProtoId: proto.Int32(0),
			Removes: &protos.Removes{
				Tile: tilerm,
			},
			Updates: &protos.Updates{
				Tile: tiles,
			},
		})
	} else {
		olds.CallRound(12, &protos.Response_S{ProtoId: proto.Int32(0),
			Removes: &protos.Removes{
				Tile: tilerm,
			},
		})
		//olds.delEyes(uid)
		//news.addEyes(uid)
		news.CallRound(12, &protos.Response_S{ProtoId: proto.Int32(0),
			Updates: &protos.Updates{
				Tile: tiles,
			},
		})
	}
}

//----------------------------------------------------------
//event
func init() {
	var BornAreaIdx, BornAreaNum int
	BornAreaIdx = 8 //TODO读库
	BornAreaNum = 0 //TODO读库

	Event.RegA("OnUserNew", func(uid int64) {
		//分配城市坐标 第一次出城再分配更好
		//City
	NEXTAREA:
		r := BornAreaIdx/AREA_ROWN + 1
		c := BornAreaIdx % AREA_ROWN
		if c == 0 {
			r--
			c = AREA_ROWN
		}
		a := Areas[r][c]

		sx := (a.col-1)*AREA_WIDTH + 1
		sy := (a.row-1)*AREA_WIDTH + 1
		for y := sy; y < sy+AREA_WIDTH; y++ {
			for x := sx; x < sx+AREA_WIDTH; x++ {
				tile := Tiles[y][x]
				block := false //TODO 阻挡
				if tile.Tp == 0 && !block {
					//if Tiles[y-1][x-1]
					tile.Tp = TILE_CITY
					tile.Val = uid
					BornAreaNum++
					if BornAreaNum >= 800 {
						BornAreaIdx = 14 //next
						BornAreaNum = 0
					}
					//saveBornAreaNum
					//Sql.Exec("update server_data set bornArea=?, bornNum=? where sid=?", BornAreaIdx,  BornAreaNum,999)

					Tiles[y][x] = tile
					Sql.Exec("replace into w_tile(x,y,tp,val) values(?,?,?,?)", tile.X, tile.Y, tile.Tp, tile.Val)
					Sql.Exec("update u_user set cityX=?, cityY=? where uid=?", tile.X, tile.Y, uid)
					//OnNewCity
					return
				}
			}
		}
		BornAreaIdx = 14 ////next
		BornAreaNum = 0
		goto NEXTAREA
	})
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {

	})
	Event.RegA("OnDisconn", func(uid int64) {
		e := Eyes.Get(uid)
		if e == nil {
			return
		}
		s := getSight(e.row, e.col)
		s.delEyes(uid)
		Eyes.Set(uid, nil)

	})

}

//RPC
func init() {

	Net.RegRPC(Rpc.View_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.View_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}

		fmt.Println("<<<View_C", ps.GetServer(), ps.GetX(), ps.GetY())

		tiles := ResetSight(ss, ps.GetServer(), ps.GetX(), ps.GetY())
		//删的前端自理
		ss.Update().Tile = tiles
		// ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
		// 	Updates: &protos.Updates{
		// 		Tile: tiles,
		// 	},
		// })
	})

	Net.RegRPC(Rpc.CityMove_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.CityMove_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<CityMove_C", ps.GetServer(), ps.GetX(), ps.GetY())
		x, y := ps.GetX(), ps.GetY()
		//TODO Item.CheckCost
		if !CheckCityLand(x, y) {
			ss.PError(protoId, 2, "CityMove_C.noMoveCity")
			return
		}
		Item.Del(uid, 2, 1, "CityMove")
		MoveCity(uid, x, y)

		Sql.Exec("update u_user set cityX=?,cityY=? where uid=?", x, y, uid)
		uu := &protos.User{
			Uid:   proto.Int64(uid),
			CityX: proto.Int32(x),
			CityY: proto.Int32(y),
		}
		updates := &protos.Updates{}
		updates.User = uu
		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Props: updates,
		})
	})

}
