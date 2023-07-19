package logic

import (
	"common/defs"
	"common/proto/client"
	"common/proto/comm"
	util2 "common/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strings"
)

func keyInputTest(u *Unit) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		World.OutMsg(defs.OpcodeCellQuit, nil)
	}
	//if inpututil.IsKeyJustPressed(ebiten.KeyC) {
	//	sendGetShop(u)
	//}
	//if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
	//	//sendbuyShop(u)
	//	fixedTimeZone(u)
	//}
	//ctrl+
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			World.OutMsg(defs.OpcodeGoToMap, &client.GoToMapReq{
				OwnerId:  0,
				MapTabId: 212,
				FriendId: 0,
			})
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			World.OutMsg(defs.OpcodeGoToMap, &client.GoToMapReq{
				OwnerId:  0,
				MapTabId: 0,
				FriendId: 0,
			})
		}
	}
	//shift+
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
	}
	//alt+
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {

	}
}

func sendGMCmd(str string) {
	//client
	args := strings.Split(str, " ")
	if len(args) == 0 {
		return
	}
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "test":
		return
	}

	//server
	gmReq := comm.GMReq{
		Action: str,
	}
	err := World.OutMsg(defs.OpcodeGM, &gmReq)
	util2.AssertTrue(err == nil, err)
}
