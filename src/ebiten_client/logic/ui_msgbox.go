package logic

import (
	ui2 "client1/logic/ebiten/ui"
	"client1/world"
	"image"
)

type UIMsgBox struct {
	ui2.BaseUI
	textMain   *ui2.TextBox
	btnConfirm *ui2.Button
	btnCancel  *ui2.Button
}

func UIShowMsgBox(text, btnText1, btnText2 string, btnClick1, btnClick2 func(b *ui2.Button)) {
	mb := NewUIMsgBox(text, btnText1, btnText2, btnClick1, btnClick2)
	ui2.ActiveUI(mb)
}

func NewUIMsgBox(text, btnText1, btnText2 string, btnClick1, btnClick2 func(b *ui2.Button)) *UIMsgBox {
	u := &UIMsgBox{}
	u.textMain = ui2.NewTextBox(image.Rect(world.ScreenWidth/2-96, 240, world.ScreenWidth/2+96, 300))
	u.btnConfirm = ui2.NewButton(image.Rect(world.ScreenWidth/2-64, 320, world.ScreenWidth/2-16, 336), "confirm")
	u.btnCancel = ui2.NewButton(image.Rect(world.ScreenWidth/2+16, 320, world.ScreenWidth/2+64, 336), "cancel")
	u.AddChild(u.textMain)
	u.AddChild(u.btnConfirm)
	u.AddChild(u.btnCancel)

	u.textMain.Text = text
	u.btnConfirm.Text = btnText1
	u.btnCancel.Text = btnText2
	u.btnConfirm.SetOnClick(func(b *ui2.Button) {
		if btnClick1 != nil {
			btnClick1(b)
		}
		ui2.CloseUI(u)
	})
	u.btnCancel.SetOnClick(btnClick2)
	u.btnCancel.SetOnClick(func(b *ui2.Button) {
		if btnClick2 != nil {
			btnClick2(b)
		}
		ui2.CloseUI(u)
	})
	return u
}
