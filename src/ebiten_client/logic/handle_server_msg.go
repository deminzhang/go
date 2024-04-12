package logic

import (
	"client0/logic/ebiten/ui"
	"common/defs"
	"common/proto/client"
	"common/proto/comm"
	"common/tlog"
	"common/util"
	"common/vec"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"os"
	"time"
)

type handleServerMsgFn func(player *Unit, msg []byte)

var handleServerMsgMap = map[uint16]handleServerMsgFn{}

func handleFrameUpdate(player *Unit, msg []byte) {
	var resp = client.StatusFrameResp{}
	proto.Unmarshal(msg, &resp)

	for _, frame := range resp.Frames {
		object := World.GetPlayer(frame.Head.ObjectId)
		if object == nil {
			continue
		}
		//logger.Debug("player:%d recv frame:%d", object.id, frame.Head.SeqId)
		if frame.Head.ObjectId == World.playerId {
			World.lastFrameSeq = frame.Head.SeqId
		} else {
			if frame.MoveFrame != nil {
				if frame.MoveFrame.Pos == nil {
					frame.MoveFrame.Pos = &comm.Vec3F{}
				}
				if frame.MoveFrame.Rot == nil {
					frame.MoveFrame.Rot = &comm.Vec3F{}
				}
				if frame.MoveFrame.Velocity == nil {
					frame.MoveFrame.Velocity = &comm.Vec3F{}
				}
				object.ApplyInput(frame)
			}
		}
	}
}

func handleFrameReset(plr *Unit, msg []byte) {
	var resp = client.StatusFrameResetResp{}
	proto.Unmarshal(msg, &resp)

	p := World.GetPlayer(resp.PlayerId)
	if p == nil {
		return
	}
	tlog.Debugf("frameReset %d %v", resp.SeqId, resp.MoveFrame)
	p.pos = *resp.MoveFrame.Pos
	p.rot = *resp.MoveFrame.Rot
	p.velocity = *resp.MoveFrame.Velocity
	p.frameComponent.ResetQueue(int(resp.SeqId))
}

func handlePlayerEnterMap(plr *Unit, msg []byte) {
	tlog.Info("handlePlayerEnterMap", plr)
	var resp client.PlayerEnterMapResp
	proto.Unmarshal(msg, &resp)
	plr.unitType = 1 //defs.UnitPlayer
	p := resp.Role
	if p.Pos != nil {
		plr.pos = *p.Pos
	}
	if p.Velocity != nil {
		plr.velocity = *p.Velocity
	}
	if p.Rot != nil {
		plr.rot = *p.Rot
	}
	World.sceneTabId = resp.TabId
	World.sceneInstId = resp.InstId
	World.sceneOwnerId = resp.OwnerId
	World.sceneFriendId = resp.FriendId
	World.sceneHouseId = resp.HouseId

	plr.frameComponent.ResetQueue(int(p.InitSeqId))
	plr.bInMap = true
	World.rpcNtp()
	World.AddMe(plr)
	UIHideLogin()
	UIShowChat()
	UIChatLog("PlayerEnterMap:%v", plr.pos)
	UIChatLog("SceneOpens:%v", resp.Opens)
	fmt.Println("SceneOpens:", resp.Opens)
	UIChatLog("BuildOpens:%v", resp.BuildOpens)
	fmt.Println("BuildOpens:", resp.BuildOpens)
}
func handleSceneOpenUpdate(player *Unit, msg []byte) {
	var resp client.SceneOpenUpdateResp
	proto.Unmarshal(msg, &resp)
	tlog.Info("handleSceneOpenUpdate", resp)
}
func handleBuildOpenUpdate(player *Unit, msg []byte) {
	var resp client.BuildOpenUpdateResp
	proto.Unmarshal(msg, &resp)
	tlog.Info("handleBuildOpenUpdate", resp.BuildOpens)
}
func handleEmotionUpdate(player *Unit, msg []byte) {
	var resp client.EmotionUpdateResp
	proto.Unmarshal(msg, &resp)
	tlog.Info("handleEmotionUpdate", resp.Opens)
}

func handleObjectSpawn(player *Unit, msg []byte) {
	var req client.SpawnUnitsResp
	proto.Unmarshal(msg, &req)

	for _, obj := range req.Info {
		if obj.Id != World.playerId {

			fmt.Println("spawn object id", obj.Id, obj.UnitType, obj.Pos)
			p := NewUnit(obj.Id)
			p.unitType = obj.UnitType
			if obj.Pos != nil {
				p.pos = *obj.Pos
			}

			if obj.Velocity != nil {
				p.velocity = *obj.Velocity
			}

			if obj.Rot != nil {
				p.rot = *obj.Rot
			}
			p.frameComponent.ResetQueue(int(obj.InitSeqId))
			p.bInMap = true
			World.AddPlayer(p)
		}
	}
}

func handleChangeMapReady(player *Unit, msg []byte) {
	tlog.Info("handleChangeMapReady")
	var req client.ChangeMapResp
	proto.Unmarshal(msg, &req)
	//req.MapTabId
	//req.From

}

func handlePlayerLeave(player *Unit, msg []byte) {
	tlog.Info("handleCellLeave")

	World.players = map[int64]*Unit{}

	//ui.AddUI("reLogin", NewUIMsgBox("已退出! ",
	//	"确认", "退出",
	//	func(b *ui.Button) {
	//		//TODO
	//		//World.Self().rpcPlayerInfo()
	//		World.Self().rpcJoinBattle()
	//	}, func(b *ui.Button) {
	//		//World.bStart = false
	//		//World.self = nil
	//		os.Exit(0)
	//	}))
}

func handleNtp(player *Unit, msg []byte) {
	var resp comm.NtpResp
	err := proto.Unmarshal(msg, &resp)
	fmt.Println(err)
}

func handleCellUnitLeave(player *Unit, msg []byte) {
	var resp client.UnitLeaveResp
	proto.Unmarshal(msg, &resp)
	if resp.Id != World.playerId {
		World.RemovePlayer(resp.Id)
	} else {
		//World.bStart = false
		World.RemovePlayer(resp.Id)
	}
}

func handleUnitMove(player *Unit, msg []byte) {
	var resp client.UnitMoveResp
	proto.Unmarshal(msg, &resp)
	unit := World.GetPlayer(resp.Id)
	if unit == nil {
		return
	}
	from := vec.Vector3{}
	target := vec.Vector3{}
	from.SetProtoVec3f(&unit.pos)
	target.SetProtoVec3f(resp.Target)
	vel := vec.SubV3(from, target)
	vel.ScaleToLength(float32(resp.Speed))
	unit.velocity = *vel.ToProtoVec3f()
	unit.rot = *vel.ToProtoVec3f()
	unit.target = target
}

func handleUnitStop(player *Unit, msg []byte) {
	var resp client.UnitStopResp
	proto.Unmarshal(msg, &resp)
	unit := World.GetPlayer(resp.Id)
	if unit == nil {
		return
	}
	unit.pos = *resp.Pos
	unit.velocity = comm.Vec3F{}

}

func handleUnitLeave(player *Unit, msg []byte) {
	var resp client.UnitLeaveTowerAoiResp
	proto.Unmarshal(msg, &resp)
	for _, id := range resp.Ids {
		World.RemovePlayer(id)
	}
}

func handleWeatherResp(player *Unit, msg []byte) {
	var resp comm.Weather
	proto.Unmarshal(msg, &resp)
	fmt.Println(resp.Weather)
	World.weather = resp.Weather
}

func handleGMResp(player *Unit, msg []byte) {
	var resp comm.GMResp
	proto.Unmarshal(msg, &resp)
	fmt.Println(resp.Result)
	if uiChat != nil {
		UIChatLog("GMResp:" + resp.Result)
	}
}

func handleDropItemList(player *Unit, msg []byte) {
	var resp client.DropItemListResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("DropItemList", resp.List)
}

func handleDropItemRemove(player *Unit, msg []byte) {
	var resp client.DropItemRemoveResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("DropItemRemove", resp.Id, resp.PlayerId)
}
func handleModifyVoxelSpanList(player *Unit, msg []byte) {
	var resp client.ModifyVoxelSpanListResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("ModifyVoxelSpanListResp", resp.Action, resp.Spans)

}

func handleStrengthResp(player *Unit, msg []byte) {
	var resp client.StrengthChangeResp
	proto.Unmarshal(msg, &resp)
	World.strength[0] = resp.Strength.Strength
	World.strength[1] = resp.Strength.StrengthMax
}

func handlePlayerLoginData(player *Unit, msg []byte) {
	var resp client.PlayerLoginDataResp
	proto.Unmarshal(msg, &resp)
	fmt.Println(resp)
	if resp.Strength != nil {
		World.strength[0] = resp.Strength.Strength
		World.strength[1] = resp.Strength.StrengthMax
	}
	//if resp.Weather != nil {
	//	World.weather = resp.Weather.Weather
	//	World.season = season2str[resp.Weather.Season]
	//}
}

func handlePlayerJoinFinished(player *Unit, msg []byte) {
	var resp client.PlayerJoinFinishResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handlePlayerJoinFinished", resp.MapTabId, resp.MapFriendId)
}

//func handlePlayerShopData(player *Unit, msg []byte) {
//	var resp client.ShopSyncResp
//	proto.Unmarshal(msg, &resp)
//	fmt.Println(resp)
//}

func handleJoinMapInvite(player *Unit, msg []byte) {
	var resp client.JoinMapInviteResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handleJoinMapInvite", resp)
	str := fmt.Sprintf("%s邀请你加入他所在场景", resp.PlayerName)
	UIShowMsgBox(str, "同意", "无视", func(b *ui.Button) {
		World.OutMsg(defs.OpcodeJoinMap, &client.JoinMapReq{
			MapInst: resp.MapInst,
			HouseId: resp.HouseId,
		})
	}, func(b *ui.Button) {
	})
}
func handleHouseRegionList(player *Unit, msg []byte) {
	var resp client.HouseRegionResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handleHouseRegionList", resp.MapTabId, resp.MapFriendId)
	for _, region := range resp.RegionList {
		fmt.Println("RegionList", region.Id)
	}
	UIChatLog("MapTabId:%d MirrorId: %d Region数 %d", resp.MapTabId, resp.MapFriendId, len(resp.RegionList))
}
func handlePlayerJumpPos(player *Unit, msg []byte) {
	var resp client.PlayerJumpPosResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handlePlayerJumpPos", resp)
	unit := World.GetPlayer(resp.Id)
	if unit == nil {
		return
	}
	unit.pos = *resp.Pos
	unit.rot = *resp.Rot
}

func handlePlantList(player *Unit, msg []byte) {
	var resp client.PlantListResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handlePlantList", len(resp.List))
	UIChatLog("MapTabId:%d MirrorId: %d 植物数 %d", resp.MapTabId, resp.MapFriendId, len(resp.List))
}
func handleDropList(player *Unit, msg []byte) {
	var resp client.DropItemListResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handlePlantList", len(resp.List))
	UIChatLog("MapTabId:%d MirrorId: %d 掉落数 %d", resp.MapTabId, resp.MapFriendId, len(resp.List))
}
func handleBuildingList(player *Unit, msg []byte) {
	var resp client.BuildingListResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("handlePlantList", len(resp.List))
	UIChatLog("MapTabId:%d MirrorId: %d 建筑数 %d", resp.MapTabId, resp.MapFriendId, len(resp.List))
}
func handleDropItemAdd(player *Unit, msg []byte) {
	var resp client.DropItemAddResp
	proto.Unmarshal(msg, &resp)
	fmt.Println("DropItemAdd", len(resp.List))
	UIChatLog("MapTabId:%d MirrorId: %d 新掉落 %d", resp.MapTabId, resp.MapFriendId, len(resp.List))
}

func handleBuildingRemove(plr *Unit, msg []byte) {
	var resp client.BuildingRemoveResp
	proto.Unmarshal(msg, &resp)

	UIChatLog("建筑删 Id:%d MapTabId:%d MirrorId: %d ", resp.Id, resp.MapTabId, resp.MapFriendId)
}
func handleBlueprintList(plr *Unit, msg []byte) {
	var resp client.BlueprintListResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图刷%d:%v", resp.Tp, len(resp.List))
	fmt.Printf("蓝图刷%d:%v", resp.Tp, resp.List)
}
func handleBlueprintSave(plr *Unit, msg []byte) {
	var resp client.BlueprintSaveResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图存resp")
	fmt.Printf("蓝图存:%v", resp.Blueprint)
}
func handleBlueprintDel(plr *Unit, msg []byte) {
	var resp client.BlueprintDelResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图删:%d", resp.Id)
}
func handleBlueprintPub(plr *Unit, msg []byte) {
	var resp client.BlueprintPubResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图发:%v", resp.Blueprint)
}
func handleBlueprintUnPub(plr *Unit, msg []byte) {
	var resp client.BlueprintUnPubResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图取消发布:%d", resp.Id)
}
func handleBlueprintSupportList(plr *Unit, msg []byte) {
	var resp client.BlueprintSupportListResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图赞列:%v", resp.Ids)
}
func handleBlueprintSupport(plr *Unit, msg []byte) {
	var resp client.BlueprintSupportResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图赞:%d %v %v", resp.Id, resp.Cancel, resp.Supports)
}
func handleBlueprintFavorite(plr *Unit, msg []byte) {
	var resp client.BlueprintFavoriteResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图收:%d %v %v", resp.Id, resp.Cancel, resp.Favorites)
}
func handleBlueprintSetPasswd(plr *Unit, msg []byte) {
	var resp client.BlueprintSetPasswdResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图密:%d %v %s", resp.Id, resp.SwPass, resp.Passwd)
}
func handleBlueprintSearch(plr *Unit, msg []byte) {
	var resp client.BlueprintSearchResp
	proto.Unmarshal(msg, &resp)
	UIChatLog("蓝图搜: %d %d", resp.Offset, len(resp.List))
	fmt.Printf("蓝图搜: %d \n%v", resp.Offset, resp.List)
}

func handleIllustratedBookSumUpdate(plr *Unit, msg []byte) {
	fmt.Println("handleIllustratedBookSumUpdate")
	var resp client.IllustratedSumResp
	proto.Unmarshal(msg, &resp)
	fmt.Printf("图鉴Sum: \n%v", resp.Sum)
}
func handleIllustratedBookList(plr *Unit, msg []byte) {
	fmt.Println("handleIllustratedBookList")
	var resp client.IllustratedBookListResp
	proto.Unmarshal(msg, &resp)
	fmt.Printf("图鉴:  \n%v", resp.Item)
}
func handleIllustratedBookRead(plr *Unit, msg []byte) {
	fmt.Println("handleIllustratedBookRead")
}
func handleIllustratedBookSumGetReward(plr *Unit, msg []byte) {
	fmt.Println("handleIllustratedBookSumGetReward")
	//var resp client.IllustratedBookSumGetRewardResp
	//proto.Unmarshal(msg, &resp)
}

// ------------------------------------------
func handlePlayerLogin(plr *Unit, msg []byte) {
	var resp comm.PlayerLoginResp
	err := proto.Unmarshal(msg, &resp)
	util.AssertTrue(err == nil, err)
	tlog.Info("handlePlayerLogin", resp)
	tlog.Info("handlePlayerLogin.EmotionIds", resp.EmotionIds)
	if resp.Homestead == 0 { //未选家
		//homeMaps := []int32{1, 2} //TODO 可选家园岛地图@map.xlsx
		err := World.OutMsg(defs.OpcodeSelectHome, &client.PlayerSelectHomeReq{
			Homestead: 1,
			//HomeIsland: homeMaps[rand.Intn(2)],
			Area: 1,
		})
		util.AssertTrue(err == nil, err)
	}
	if resp.NeedRebuilding {
		World.rpcRebuilding()
	} else {
		World.rpcJoinBattle()
	}
}
func handleLogout(plr *Unit, msg []byte) {
	tlog.Info("handleLogout")
	os.Exit(0)
}
func handleEmpty(plr *Unit, msg []byte) {
	fmt.Println("handleEmpty")
}
func init() {
	rand.Seed(time.Now().UnixNano())
	handleServerMsgMap[defs.OpcodeFrameUpdateResp] = handleFrameUpdate
	handleServerMsgMap[defs.OpcodeFrameResetResp] = handleFrameReset
	handleServerMsgMap[defs.OpcodePlayerEnterMapResp] = handlePlayerEnterMap
	handleServerMsgMap[defs.OpcodeSceneOpenUpdate] = handleSceneOpenUpdate
	handleServerMsgMap[defs.OpcodeBuildOpenUpdateResp] = handleBuildOpenUpdate
	handleServerMsgMap[defs.OpcodeEmotionUpdateResp] = handleEmotionUpdate
	handleServerMsgMap[defs.OpcodeObjectSpawnResp] = handleObjectSpawn
	handleServerMsgMap[defs.OpcodeNtpResp] = handleNtp
	handleServerMsgMap[defs.OpcodeCellQuitResp] = handleCellUnitLeave
	handleServerMsgMap[defs.OpcodeCellLeaveResp] = handlePlayerLeave
	handleServerMsgMap[defs.OpcodeCellChangeResp] = handleChangeMapReady
	handleServerMsgMap[defs.OpcodeUnitMoveResp] = handleUnitMove
	handleServerMsgMap[defs.OpcodeUnitStopResp] = handleUnitStop
	handleServerMsgMap[defs.OpcodeAoiLeaveViewResp] = handleUnitLeave
	handleServerMsgMap[defs.OpcodeCellWeatherResp] = handleWeatherResp
	handleServerMsgMap[defs.OpcodeGMResp] = handleGMResp
	handleServerMsgMap[defs.OpcodeDropItemListResp] = handleDropItemList
	handleServerMsgMap[defs.OpcodeDropItemAddResp] = handleDropItemAdd
	handleServerMsgMap[defs.OpcodeDropItemRemoveResp] = handleDropItemRemove
	handleServerMsgMap[defs.OpcodeModifyVoxelSpanResp] = handleModifyVoxelSpanList
	handleServerMsgMap[defs.OpcodeStrengthUpdateResp] = handleStrengthResp
	//handleServerMsgMap[defs.OpcodePlayerLoginDataResp] = handlePlayerLoginData
	//handleServerMsgMap[defs.OpcodePlayerShopSyncResp] = handlePlayerShopData
	handleServerMsgMap[defs.OpcodePlayerJoinFinishedResp] = handlePlayerJoinFinished
	handleServerMsgMap[defs.OpcodeJoinMapInviteResp] = handleJoinMapInvite
	handleServerMsgMap[defs.OpcodeHouseRegionListResp] = handleHouseRegionList
	handleServerMsgMap[defs.OpcodePlayerJumpPosResp] = handlePlayerJumpPos
	handleServerMsgMap[defs.OpcodePlantListResp] = handlePlantList
	handleServerMsgMap[defs.OpcodeDropItemListResp] = handleDropList
	handleServerMsgMap[defs.OpCodeBuildingListResp] = handleBuildingList

	handleServerMsgMap[defs.OpCodeBuildingBuildResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildingSaveDIYResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildingUpdateResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildingRemoveResp] = handleBuildingRemove
	handleServerMsgMap[defs.OpCodeBuildingMoveResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildCancelResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildDonateResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildSpeedUpResp] = handleEmpty
	handleServerMsgMap[defs.OpCodeBuildCompleteResp] = handleEmpty
	handleServerMsgMap[defs.OpcodeBlueprintListResp] = handleBlueprintList
	handleServerMsgMap[defs.OpcodeBlueprintSaveResp] = handleBlueprintSave
	handleServerMsgMap[defs.OpcodeBlueprintDelResp] = handleBlueprintDel
	handleServerMsgMap[defs.OpcodeBlueprintPubResp] = handleBlueprintPub
	handleServerMsgMap[defs.OpcodeBlueprintUnPubResp] = handleBlueprintUnPub
	handleServerMsgMap[defs.OpcodeBlueprintSupportListResp] = handleBlueprintSupportList
	handleServerMsgMap[defs.OpcodeBlueprintSupportResp] = handleBlueprintSupport
	handleServerMsgMap[defs.OpcodeBlueprintFavoriteResp] = handleBlueprintFavorite
	handleServerMsgMap[defs.OpcodeBlueprintSetPasswdResp] = handleBlueprintSetPasswd
	handleServerMsgMap[defs.OpcodeBlueprintSearchResp] = handleBlueprintSearch
	handleServerMsgMap[defs.OpcodePlayerLoginResp] = handlePlayerLogin
	handleServerMsgMap[defs.OpcodeTryLogoutResp] = handleLogout
	handleServerMsgMap[defs.OpcodeIllustratedBookSumUpdate] = handleIllustratedBookSumUpdate
	handleServerMsgMap[defs.OpcodeIllustratedBookListResp] = handleIllustratedBookList
	handleServerMsgMap[defs.OpcodeIllustratedBookReadResp] = handleIllustratedBookRead
	handleServerMsgMap[defs.OpcodeIllustratedBookSumGetRewardResp] = handleIllustratedBookSumGetReward

}
