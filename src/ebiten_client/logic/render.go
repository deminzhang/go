package logic

import (
	"client0/logic/asset"
	"client0/logic/ebiten/ui"
	"client0/world"
	"common/defs"
	"common/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
	"math"
)

type TextLine struct {
	Content string
}
type TextComponent struct {
	Lines      []TextLine
	Color      color.NRGBA
	LineHeight int
	Font       font.Face
	IsCentered bool
}

var (
	lanaPixel font.Face
	comText   *TextComponent
	offsetX   float64 = 300
	offsetZ   float64 = 50
)

func init() {
	var err error
	lanaPixel, err = asset.LoadFont("font/lana_pixel.ttf", nil)
	if err != nil {
		return
	}
	comText = &TextComponent{
		Color:      color.NRGBA{50, 220, 100, 255},
		Font:       lanaPixel,
		LineHeight: 10,
	}
}

func (this *game) Draw(screen *ebiten.Image) {
	ui.Draw(screen)
	if !this.bStart {
		return
	}

	var sceneW, sceneH float64 = 300, 300 //TODO 从scene上读
	ebitenutil.DrawRect(screen, offsetX, offsetZ, sceneW, sceneH, color.Gray{Y: 128})
	for _, plr := range this.players {
		this.DrawUnit(plr, screen)
	}
	text.Draw(screen, fmt.Sprintf("w%d h%d", int(sceneW), int(sceneH)),
		lanaPixel, int(offsetX), int(sceneH+offsetZ+20), comText.Color)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("playerId: %d\n"+
			"scene: %d x: %.2f z: %.2f ownerId: %d\n"+
			"mapInstId: %d houseId: %d friendId: %d \n"+
			"weather: %d - %d - %d -%d \n"+
			"strength : %d / %d \n%s"+
			"confirmFrameSeqId: %d",
			this.playerId,
			this.sceneTabId, this.self.pos.X, this.self.pos.Z, this.sceneOwnerId,
			this.sceneInstId, this.sceneHouseId, this.sceneFriendId,
			this.weather[0], this.weather[1], this.weather[2], this.weather[3],
			this.strength[0], this.strength[1], this.season,
			World.lastFrameSeq,
		))
}

func (this *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return world.ScreenWidth, world.ScreenHeight
}

func (this *game) playerSprite() *ebiten.Image {
	img := this.spriteMap["player"]
	if img == nil {
		rawImg, err := asset.LoadImage("sprite/yellowball.png")
		util.AssertTrue(err == nil, "image decode :%v", err)

		img = ebiten.NewImage(32, 32)
		img.DrawImage(ebiten.NewImageFromImage(rawImg), nil)

		this.spriteMap["player"] = img
		return img
	}
	return img
}
func (this *game) npcSprite() *ebiten.Image {
	img := this.spriteMap["npc"]
	if img == nil {
		rawImg, err := asset.LoadImage("sprite/player.png")
		util.AssertTrue(err == nil, "image decode :%v", err)

		img = ebiten.NewImage(32, 32)
		img.DrawImage(ebiten.NewImageFromImage(rawImg), nil)

		this.spriteMap["npc"] = img
		return img
	}
	return img
}

func (this *game) DrawUnit(unit *Unit, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	size := 1.0 / 4
	op.GeoM.Scale(size, size)
	op.GeoM.Translate(float64(unit.pos.X)+offsetX, float64(unit.pos.Z)+offsetZ)

	switch unit.unitType {
	case defs.UnitPlayer:
		screen.DrawImage(this.playerSprite(), op)
	default:
		screen.DrawImage(this.npcSprite(), op)
	}

	//draw playerId,posX,posY
	if World.renderShowId {
		var lines []TextLine
		lines = append(lines, TextLine{Content: fmt.Sprintf("p%d", unit.id)})
		//lines = append(lines, TextLine{Content: fmt.Sprintf("x%.2f", unit.pos.X)})
		//lines = append(lines, TextLine{Content: fmt.Sprintf("y%.2f", unit.pos.Z)})
		var size float64
		//size = 64
		for i, line := range lines {
			y := (size / 2) - float64((comText.LineHeight*len(comText.Lines))/2) + float64((i+1)*comText.LineHeight)

			text.Draw(screen, line.Content, lanaPixel, int(unit.pos.X+float32(offsetX)), int(unit.pos.Z+float32(y)+float32(offsetZ)), comText.Color)
		}
	}
}

func (this *Unit) collidesWithWall() bool {
	x := int(math.Ceil(float64(this.pos.X)))
	y := int(math.Ceil(float64(this.pos.Z)))

	bx := int(math.Ceil(float64(this.boundingBox.X)))
	by := int(math.Ceil(float64(this.boundingBox.Z)))

	if x < bx || y < by || x > world.ScreenWidth-bx || y > world.ScreenHeight-by {
		return true
	}
	return false
}
