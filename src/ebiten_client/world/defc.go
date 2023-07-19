package world

import (
	. "common/fix64"
	"math"
)

const (
	TPSRate = 10
	Speed   = float32(50.0)
)

var Inv = NewFix64(1000.0 / float32(TPSRate))
var FixSpeed = NewFix64(0.08)
var FixFlySpeedOne = NewFix64(0.2)
var FixFlySpeedTwo = NewFix64(1)

var TPSInterval = int(math.Ceil(float64(1) / float64(TPSRate)))
var ReportInputInterval = TPSInterval/2 + 1
var LagSendInterval = 0

const (
	StatusFrameTypeMove = 1
)

const (
	KeyNone    = 0
	KeyLeft    = 1
	KeyRight   = 2
	KeyUp      = 4
	UpperLeft  = 5 //左上
	UpperRight = 6 //右上
	KeyDown    = 8
	LowerLeft  = 9  //左下
	LowerRight = 10 //右下
	KeyK       = 16
	KeyJ       = 32
	KeyL       = 64
)

const (
	ScreenWidth     = 640
	ScreenHeight    = 480
	GridSize        = 5
	PlayerBoundingX = 32
	PlayerBoundingY = 32

	xNumInScreen = ScreenWidth / GridSize
	yNumInScreen = ScreenHeight / GridSize
)
