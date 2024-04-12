package world

import (
	"math"
)

const (
	TPSRate = 10
	Speed   = float32(50.0)
)

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
	ScreenWidth     = 800
	ScreenHeight    = 600
	GridSize        = 5
	PlayerBoundingX = 32
	PlayerBoundingY = 32

	xNumInScreen = ScreenWidth / GridSize
	yNumInScreen = ScreenHeight / GridSize
)
