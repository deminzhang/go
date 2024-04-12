package logic

import (
	"client0/logic/ebiten/ui"
	utilc "client0/util"
	"client0/world"
	"common/defs"
	"common/proto/client"
	"common/proto/comm"
	"common/util"
	"common/vec"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wonderivan/logger"
	"math"
)

const (
	UnitBase     = 0 //空姬
	UnitPlayer   = 1 //玩家
	UnitNpc      = 2 //NPC
	UnitBullet   = 3 //子弹
	UnitLaser    = 4 //闪光(后端无.仅通知前端)
	UnitBuilding = 5
)

type ProtoMsg struct {
	op  uint16
	msg []byte
}

type Unit struct {
	id             int64
	unitType       int32
	netSeqId       int32
	frameComponent *world.FrameComponent
	pos            comm.Vec3F
	rot            comm.Vec3F
	velocity       comm.Vec3F
	target         vec.Vector3
	lastSendDt     int
	accuDt         int
	bInMap         bool
	boundingBox    comm.Vec3F
	lastPressKey   int
	author         string
}

func NewUnit(id int64) *Unit {
	return &Unit{
		id:             id,
		frameComponent: world.NewFrameComponent(world.WithAssert()),
		boundingBox:    comm.Vec3F{X: world.PlayerBoundingX, Y: world.PlayerBoundingY, Z: 32},
	}
}

func (this *Unit) update(now int64, dt int) {
	if World.Self() == this {
		this.meUpdate(now, dt)
	} else {
		switch this.unitType {
		//case UnitPlayer:
		//	this.otherFrameUpdate(now, dt)
		//case UnitNpc:
		//	this.otherUpdate(now, dt)
		//default:
		//	this.otherUpdate(now, dt)
		}
		this.otherFrameUpdate(now, dt)
	}
}

func (this *Unit) otherUpdate(now int64, dt int) {
	if !this.bInMap {
		return
	}
	this.pos.X = this.pos.X + float32(dt)*this.velocity.X/1000
	this.pos.Z = this.pos.Z + float32(dt)*this.velocity.Z/1000
	this.repairPosBounding()
}

func (this *Unit) otherFrameUpdate(now int64, dt int) {
	if !this.bInMap {
		return
	}

	this.accuDt += dt
	for {
		if this.frameComponent.IsEmpty() {
			this.accuDt = 0
			break
		}
		frame := this.frameComponent.HeadElement().(*FrameData)
		//logger.Warn(" other update accu dt :%d dt:%d move:%+v", this.accuDt, dt, *frame.MoveFrame)
		if this.accuDt < int(frame.MoveFrame.Dt) {
			this.pos.X = frame.MoveFrame.Pos.X + float32(this.accuDt)*frame.MoveFrame.Velocity.X/1000
			this.pos.Y = frame.MoveFrame.Pos.Y + float32(this.accuDt)*frame.MoveFrame.Velocity.Y/1000
			this.pos.Z = frame.MoveFrame.Pos.Z + float32(this.accuDt)*frame.MoveFrame.Velocity.Z/1000
			this.repairPosBounding()
			break
		} else {
			dealt := frame.MoveFrame.Dt
			this.pos.X = frame.MoveFrame.Pos.X + float32(dealt)*frame.MoveFrame.Velocity.X/1000
			this.pos.Y = frame.MoveFrame.Pos.Y + float32(dealt)*frame.MoveFrame.Velocity.Y/1000
			this.pos.Z = frame.MoveFrame.Pos.Z + float32(dealt)*frame.MoveFrame.Velocity.Z/1000
			this.frameComponent.ConfirmSeqId(int(frame.Head.SeqId))
			this.accuDt -= int(dealt)
			this.repairPosBounding()
		}
	}
}

func (this *Unit) gatherKeyInput() int {
	keyInputTest(this) //测试用的快捷写这里面!!!

	key := this.lastPressKey
	if ui.Focused() {
		return key
	}
	//测试写到keyInputTest里!!!

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		key |= world.KeyLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		key |= world.KeyRight
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		key |= world.KeyDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		key |= world.KeyUp
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) || inpututil.IsKeyJustReleased(ebiten.KeyA) {
		key = key & (^world.KeyLeft)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) || inpututil.IsKeyJustReleased(ebiten.KeyD) {
		key = key & (^world.KeyRight)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) || inpututil.IsKeyJustReleased(ebiten.KeyS) {
		key = key & (^world.KeyDown)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) || inpututil.IsKeyJustReleased(ebiten.KeyW) {
		key = key & (^world.KeyUp)
	}
	if key&0b11 == 0b11 {
		key = key ^ 0b11
	}
	if key&0b1100 == 0b1100 {
		key = key ^ 0b1100
	}

	var vel = comm.Vec3F{}
	if key&world.KeyLeft == world.KeyLeft {
		if key&0b1100 > 0 {
			vel.X = -world.Speed / 1.4
			if key&world.KeyUp == world.KeyUp {
				vel.Z = -world.Speed / 1.4
			}
			if key&world.KeyDown == world.KeyDown {
				vel.Z = world.Speed / 1.4
			}
		} else {
			vel.X = -world.Speed
		}
	} else if key&world.KeyRight == world.KeyRight {
		if key&0b1100 > 0 {
			vel.X = world.Speed / 1.4
			if key&world.KeyUp == world.KeyUp {
				vel.Z = -world.Speed / 1.4
			}
			if key&world.KeyDown == world.KeyDown {
				vel.Z = world.Speed / 1.4
			}
		} else {
			vel.X = world.Speed
		}
	} else if key&world.KeyUp == world.KeyUp {
		vel.Z = -world.Speed
	} else if key&world.KeyDown == world.KeyDown {
		vel.Z = world.Speed
	}
	this.velocity = vel

	return key
}

func (this *Unit) meUpdate(now int64, dt int) {
	if this.bInMap {
		this.genFrame(dt)
		this.lastPressKey = this.gatherKeyInput()
	}
}

func (this *Unit) genFrame(dt int) {
	if this.lastPressKey != world.KeyNone {
		frameComponent := this.frameComponent
		pos := this.pos
		vel := this.velocity

		v3 := vec.Vector3{}
		v3.X = pos.X + vel.X*float32(dt)/1000
		v3.Z = pos.Z + vel.Z*float32(dt)/1000
		v3.Y = pos.Y

		//y, _ := World.scene.CheckBlocked(v3, 10)
		//if y == -1 {
		//	fmt.Println("mePlayerMove Block")
		//	return
		//} else {
		this.pos.X = v3.X
		this.pos.Z = v3.Z
		//	this.pos.Y = World.scene.ToRealHeight(y)
		//}

		frame := comm.StatusFrame{
			Head: &comm.FrameHead{
				SeqId:      int32(frameComponent.TailSeqId()),
				ObjectId:   this.id,
				ServerTick: World.serviceTick(),
				TypeId:     world.StatusFrameTypeMove,
			},

			MoveFrame: &comm.MoveFrame{
				Pos:      &pos,
				Velocity: &vel,
				Rot:      &vel, //comm.Vec3F{},
				Dt:       int32(dt),
			},
		}
		this.frameComponent.EnQueue(&FrameData{
			StatusFrame: &frame,
			FrameType:   defs.FrameTypeClient,
		})
	}
	this.lastSendDt += dt

	if this.lastSendDt >= world.ReportInputInterval {
		this.sendFrames()
		this.lastSendDt = this.lastSendDt % world.ReportInputInterval
	}
}

func (this *Unit) sendFrames() {
	var frames []*comm.StatusFrame
	this.frameComponent.TraceStart(
		this.frameComponent.NextSendSeqId(),
		func(element utilc.IFrameElement) bool {
			data := element.(*FrameData)
			fmt.Println(fmt.Sprintf("now send seqId :%d", data.GetSeqId()))
			frames = append(frames, data.StatusFrame)
			return false
		},
	)

	size := this.frameComponent.TailSeqId() - this.frameComponent.NextSendSeqId()
	this.frameComponent.SetNextSendSeqId(this.frameComponent.TailSeqId())
	util.AssertTrue(this.frameComponent.NextSendSeqId() == this.frameComponent.TailSeqId(), "")
	if len(frames) > 0 {
		util.AssertTrue(size == len(frames), "sendFrames size error")
		err := World.OutMsg(defs.OpcodeFrameInput, &client.StatusFrameInput{Frames: frames})
		//this.lagSendMsg(defs.OpcodeFrameInput, &client.StatusFrameInput{Frames: frames})
		logger.Info("send frames :%+v", frames)
		if err != nil {
			logger.Info("err:%v", err)
		}
	}
}

func (this *Unit) DailBack(frame *comm.StatusFrame) {
	frameSeqId := int(frame.Head.SeqId)
	logger.Warn("this :%d dail back seq id :%d", this.id, frame.Head.SeqId)
	if this.frameComponent.DailBack(frameSeqId) {
		oldData := this.GetFrameData(frameSeqId)
		logger.Info("this:%d replace seq old:+%v new :%+v", this.id, oldData, frame)
		this.frameComponent.Replace(&FrameData{
			StatusFrame: frame,
			FrameType:   defs.FrameTypeServer,
		})

		this.pos = *frame.MoveFrame.Pos
		this.rot = *frame.MoveFrame.Rot
		this.velocity = *frame.MoveFrame.Velocity

		this.frameComponent.TraceStart(
			frameSeqId,
			func(element utilc.IFrameElement) bool {
				elementFrame := element.(*FrameData)
				this.recalcFrame(elementFrame)

				return false
			},
		)
		this.frameComponent.ConfirmSeqId(frameSeqId)
	}
}

func (this *Unit) recalcFrame(frame *FrameData) {
	pos := this.pos
	vel := *frame.MoveFrame.Velocity
	rot := *frame.MoveFrame.Rot

	frame.MoveFrame.Pos = utilc.CopyVec3f(&pos)
	pos.X = pos.X + vel.X*float32(frame.MoveFrame.Dt)/1000
	pos.Y = pos.Y + vel.Y*float32(frame.MoveFrame.Dt)/1000
	pos.Z = pos.Z + vel.Z*float32(frame.MoveFrame.Dt)/1000

	logger.Info("repair seq id :%d frame pre pos:%+v vel:%+v new pos:%+v",
		frame.Head.SeqId,
		pos,
		vel,
		pos,
	)
	this.pos = pos
	this.velocity = vel
	this.rot = rot
}

func (this *Unit) otherApplyInput(frame *comm.StatusFrame) {
	frameSeqId := int(frame.Head.SeqId)
	if frameSeqId == this.frameComponent.TailSeqId() {
		this.frameComponent.EnQueue(
			&FrameData{
				StatusFrame: frame,
				FrameType:   defs.FrameTypeServer,
			},
		)
		//logger.Info("other this:%d recv seq id :%d and frame:%+v", this.id, frameSeqId, *frame)
	} else {
		data := this.GetFrameData(frameSeqId)
		if int(data.Head.SeqId) != frameSeqId {
			logger.Debug("other this:%d recv un match seq id:%d", this.id, frameSeqId)
			return
		}

		if this.frameComponent.DailBack(frameSeqId) {
			this.frameComponent.Replace(&FrameData{
				StatusFrame: frame,
				FrameType:   defs.FrameTypeServer,
			})

			this.pos = *frame.MoveFrame.Pos
			this.rot = *frame.MoveFrame.Rot
			this.velocity = *frame.MoveFrame.Velocity

			this.frameComponent.TraceStart(frameSeqId,
				func(element utilc.IFrameElement) bool {
					this.recalcFrame(element.(*FrameData))
					return false
				},
			)
		}
	}
}

func (this *Unit) GetFrameData(index int) *FrameData {
	elem := this.frameComponent.GetElement(index)
	if elem != nil {
		return elem.(*FrameData)
	}
	return nil
}
func (this *Unit) meApplyInput(frame *comm.StatusFrame) {
	if frame.Head.SeqId != int32(this.frameComponent.HeadSeqId()) {
		util.AssertTrue(false,
			"object except seqId :%d but real seqId:%d",
			this.frameComponent.HeadSeqId(),
			frame.Head.SeqId,
		)
		return
	}

	frameSeqId := int(frame.Head.SeqId)
	data := this.GetFrameData(frameSeqId)
	if data != nil && data.GetSeqId() == frameSeqId {
		//check
		if !utilc.IsEqualVec3f(frame.MoveFrame.Pos, data.MoveFrame.Pos) ||
			!utilc.IsEqualVec3f(frame.MoveFrame.Velocity, data.MoveFrame.Velocity) ||
			frame.MoveFrame.Dt != data.MoveFrame.Dt {
			this.DailBack(frame)
		} else {
			if frameSeqId >= this.frameComponent.HeadSeqId() && frameSeqId <= this.frameComponent.TailSeqId() {
				this.frameComponent.ConfirmSeqId(frameSeqId)
			} else {
				this.DailBack(frame)
			}
		}
	}
}

func (this *Unit) ApplyInput(frame *comm.StatusFrame) {
	if World.Self() != this {
		this.otherApplyInput(frame)
	} else {
		this.meApplyInput(frame)
	}
}

func (this *Unit) repairPosBounding() {
	bx := float32(math.Ceil(float64(this.boundingBox.X)))
	bz := float32(math.Ceil(float64(this.boundingBox.Z)))

	if this.pos.X < 0 {
		this.pos.X = 0
	}
	if this.pos.X > world.ScreenWidth-bx {
		this.pos.X = world.ScreenWidth - bx
	}

	if this.pos.Z < 0 {
		this.pos.Z = 0
	}

	if this.pos.Z > world.ScreenHeight-bz {
		this.pos.Z = world.ScreenHeight - bz
	}
}
