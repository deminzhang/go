package logic

import (
	"client1/logic/ebiten/ui"
	"common/defs"
	"fmt"
	"image"
)

type UIChat struct {
	ui.BaseUI
	btnChatSwitch  *ui.Button
	btnClear       *ui.Button
	textBoxLog     *ui.TextBox
	inputBoxChat   *ui.InputBox
	checkBoxShowId *ui.CheckBox
	checkBoxGM     *ui.CheckBox
	btnQuitScene   *ui.Button
	btnChatSend    *ui.Button
	showChatUI     bool
}

var uiChat *UIChat

func UIShowChat() {
	if uiChat == nil {
		uiChat = NewUIChat()
		ui.ActiveUI(uiChat)
	}
}
func UIHideChat() {
	if uiChat != nil {
		ui.CloseUI(uiChat)
		uiChat = nil
	}
}
func UIChatLog(msg string, a ...any) {
	if uiChat != nil {
		uiChat.ChatLog(msg, a...)
	}
}

func NewUIChat() *UIChat {
	p := &UIChat{showChatUI: true}
	p.textBoxLog = ui.NewTextBox(image.Rect(16, 230, 288, 421))
	p.inputBoxChat = ui.NewInputBox(image.Rect(16, 432, 288, 464))
	p.checkBoxShowId = ui.NewCheckBox(300, 385, "显示Id")
	p.checkBoxGM = ui.NewCheckBox(300, 415, "GM")
	p.btnChatSend = ui.NewButton(image.Rect(290, 432, 322, 464), "发")

	p.btnChatSwitch = ui.NewButton(image.Rect(0, 464, 32, 480), "隐")
	p.btnClear = ui.NewButton(image.Rect(32, 464, 64, 480), "清")
	p.btnQuitScene = ui.NewButton(image.Rect(64, 464, 96, 480), "退")

	p.AddChild(p.btnChatSwitch)
	p.AddChild(p.btnClear)
	p.AddChild(p.textBoxLog)
	p.AddChild(p.inputBoxChat)
	p.AddChild(p.checkBoxShowId)
	p.AddChild(p.checkBoxGM)
	p.AddChild(p.btnChatSend)
	p.AddChild(p.btnQuitScene)

	p.inputBoxChat.DefaultText = "输入信息发送.."

	p.inputBoxChat.SetOnPressEnter(func(i *ui.InputBox) {
		if !p.showChatUI {
			return
		}
		if i.Focused() {
			p.btnChatSend.Click()
		} else {
			i.SetFocused(true)
		}
	})
	p.checkBoxShowId.SetChecked(World.renderShowId)
	p.checkBoxShowId.SetOnCheckChanged(func(c *ui.CheckBox) {
		World.renderShowId = c.Checked()
	})
	p.checkBoxGM.SetOnCheckChanged(func(c *ui.CheckBox) {
		msg := "debug command"
		if c.Checked() {
			msg += " (On)"
		} else {
			msg += " (Off)"
		}
		p.textBoxLog.AppendLine(msg)
	})
	p.checkBoxGM.SetChecked(true)
	p.btnChatSend.SetOnClick(func(b *ui.Button) {
		i := p.inputBoxChat
		if i.Text != "" {
			i.SetFocused(false)
			if p.checkBoxGM.Checked() { //调试命令
				p.textBoxLog.AppendLine("GM:" + i.Text)
				fmt.Println("GM:" + i.Text)
				sendGMCmd(i.Text)
			} else { //聊天
				p.textBoxLog.AppendLine("me:" + i.Text)
				//sendChat(World.self, i.Text)
			}
			i.AppendTextHistory(i.Text)
			i.Text = ""
		} else {
			i.SetFocused(false)
			p.textBoxLog.AppendLine("no input msg")
		}
	})
	p.btnChatSwitch.SetOnClick(func(b *ui.Button) {
		p.showChatUI = !p.showChatUI
		if p.showChatUI {
			b.Text = "隐"
		} else {
			b.Text = "聊"
		}
	})
	p.btnClear.SetOnClick(func(b *ui.Button) {
		p.textBoxLog.Text = ""
	})
	p.btnQuitScene.SetOnClick(func(b *ui.Button) {
		sendQuitScene()
	})
	uiChat = p
	return p
}

func (p *UIChat) Update() {
	p.BaseUI.Update()
	p.textBoxLog.Visible = p.showChatUI
	p.inputBoxChat.Visible = p.showChatUI
	p.btnChatSend.Visible = p.showChatUI
	p.checkBoxShowId.Visible = p.showChatUI
	p.checkBoxGM.Visible = p.showChatUI
	p.btnClear.Visible = p.showChatUI
	p.btnQuitScene.Visible = p.showChatUI
}

func (p *UIChat) OnClose() {
	uiChat = nil
}

func (p *UIChat) ChatLog(msg string, a ...any) {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg, a...)
	}
	p.textBoxLog.AppendLine(msg)
}

func sendQuitScene() {
	if World.conn == nil {
		return
	}
	World.OutMsg(defs.OpcodeTryLogout, nil)
}
