package logic

import (
	"client1/logic/asset"
	"client1/logic/ebiten/ui"
	"client1/util"
	"client1/world"
	"common/defs"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wonderivan/logger"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	StateInit = iota
	StateLogin
	StateGame
)

type gameWorld struct {
	gameState      int
	playerId       int64
	self           *Unit
	players        map[int64]*Unit
	ntpClientTick  int64
	ntpServiceTick int64
	cumulateDt     int64
	bStart         bool
	spriteMap      map[string]*ebiten.Image
	lastTick       int64
	dt             int
	Die            chan struct{}
	exitTick       int64

	conn         *websocket.Conn
	netSeqId     int32
	receiveQueue *util.ConcurrentQueue[*ProtoMsg]
	lastPingTick int64

	UserAddr     string
	weather      []int32
	strength     []int64
	season       string
	renderShowId bool
	lastFrameSeq int32

	sceneTabId    int32
	sceneInstId   int64
	sceneOwnerId  int64
	sceneFriendId int64
	sceneHouseId  int64
}

var World = &gameWorld{
	players:      map[int64]*Unit{},
	receiveQueue: util.NewConcurrentQueue[*ProtoMsg](),
	spriteMap:    map[string]*ebiten.Image{},
	weather:      []int32{0, 0, 0, 0},
	strength:     []int64{0, 0},
	season:       "",
	Die:          make(chan struct{}, 1),
}

func (this *gameWorld) GetPlayer(playerId int64) *Unit {
	return this.players[playerId]
}

func (this *gameWorld) AddMe(player *Unit) {
	this.self = player
	this.players[player.id] = player
	this.playerId = player.id
}

func (this *gameWorld) AddPlayer(player *Unit) {
	this.players[player.id] = player
}

func (this *gameWorld) RemovePlayer(id int64) {
	delete(this.players, id)
}

func (this *gameWorld) Self() *Unit {
	return this.self
}

func (this *gameWorld) serviceTick() int64 {
	return this.ntpServiceTick + this.cumulateDt
}

func (this *gameWorld) otherUpdate(now int64, dt int) {
	for _, player := range this.players {
		if player != this.self {
			player.update(now, dt)
		}
	}
}

func (this *gameWorld) Update() error {
	ui.Update()
	if !this.bStart {
		return nil
	}

	this.cumulateDt = this.lastTick - this.ntpClientTick

	this.handleRcvPacket()

	nowTick := time.Now().UnixMilli()
	if this.exitTick != 0 && nowTick >= this.exitTick {
		os.Exit(1)
	}
	this.ping(nowTick)
	this.innerUpdate(nowTick, this.dt)
	this.dt = int(nowTick - this.lastTick)
	if this.dt < 0 {
		this.dt = 0
	}
	this.lastTick = nowTick
	this.cumulateDt += int64(this.dt)
	return nil
}

func (this *gameWorld) initLogger(fileName string) {
	if len(strings.TrimSpace(fileName)) > 0 {
		var p = fmt.Sprintf(`{
		"Console": {
			"level": "DEBG",
			"color": true
		},
		"File": {
			"filename": "%s",
			"level": "INFO",
			"daily": true,
			"maxlines": 1000000,
			"maxsize": 256,
			"maxdays": -1,
			"append": true,
			"permit": "0660"
		}
	}`, fileName)
		logger.SetLogger(p)
	}
}

func (this *gameWorld) ShowLogin(host, user, passwd string) *gameWorld {
	UIShowLogin(host, user, passwd, defs.HomeSteadName)

	icon16, err := asset.LoadImage("images/icon_16x16.png")
	if err != nil {
		log.Fatal("loading icon_16: %w", err)
	}
	icon32, err := asset.LoadImage("images/icon_32x32.png")
	if err != nil {
		log.Fatal("loading icon_32: %w", err)
	}
	ebiten.SetWindowIcon([]image.Image{icon32, icon16})
	ebiten.SetMaxTPS(world.TPSRate)
	ebiten.SetWindowSize(world.ScreenWidth, world.ScreenHeight)
	args := util.Args2Map()
	if val, ok := args["report"]; ok {
		iv, err := strconv.Atoi(val)
		if err == nil && (iv > world.ReportInputInterval) {
			world.ReportInputInterval = iv
		}
	}

	if val, ok := args["log"]; ok {
		this.initLogger(val)
		logger.Debug("test")
	}

	if val, ok := args["lag"]; ok {
		iv, err := strconv.Atoi(val)
		if err == nil && (iv > 0) {
			world.LagSendInterval = iv
		}
	}
	ebiten.SetWindowTitle("Home(Login)")

	return this
}

func (this *gameWorld) Enter(server string, author string) {
	//assetsDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//assetsDir := "../common/"
	//assetsDir += "/asset/terrain/voxel.bin"
	//static := &voxel.VoxelStaticVolume{}
	//err := static.Load(assetsDir)
	//util.AssertTrue(err == nil, err)
	//cell := voxel.NewVoxelVolume(static)
	//this.scene = &Scene{cell}

	this.self = NewUnit(0)
	this.self.author = author

	this.startNet(server)
	time.Sleep(time.Second * 2)

}

func (this *gameWorld) innerUpdate(now int64, dt int) {
	for _, player := range this.players {
		player.update(now, dt)
	}
}
