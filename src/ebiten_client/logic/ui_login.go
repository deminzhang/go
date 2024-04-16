package logic

import (
	"client0/logic/asset"
	ui2 "client0/logic/ebiten/ui"
	"client0/util"
	"client0/world"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
)

const (
	DevHost   = "dev.metacard.gg"
	TestHost  = "clubkoala.test.metacard.gg"
	LocalHost = "localhost"
)

func init() {
	f, _ := asset.LoadFont("font/lana_pixel.ttf", &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	ui2.SetDefaultUIFont(f)
}

type UILogin struct {
	ui2.BaseUI
	inputBoxHost *ui2.InputBox
	inputBoxUser *ui2.InputBox
	inputBoxPass *ui2.InputBox

	cbNginx    *ui2.CheckBox
	cbUseLogin *ui2.CheckBox

	btnLogin     *ui2.Button
	btnDevHost   *ui2.Button
	btnTestHost  *ui2.Button
	btnLocalhost *ui2.Button
}

var uiLogin *UILogin

func UIShowLogin(host, user, pass, gameName string) {
	if uiLogin == nil {
		uiLogin = NewUILogin(host, user, pass, gameName)
		ui2.ActiveUI(uiLogin)
	}
}
func UIHideLogin() {
	if uiLogin != nil {
		ui2.CloseUI(uiLogin)
		uiLogin = nil
	}
}

func NewUILogin(host, user, pass, gameName string) *UILogin {
	p := &UILogin{}
	p.inputBoxHost = ui2.NewInputBox(image.Rect(world.ScreenWidth/2-96, 100, world.ScreenWidth/2+96, 132))
	p.inputBoxUser = ui2.NewInputBox(image.Rect(world.ScreenWidth/2-96, 136, world.ScreenWidth/2+96, 168))
	p.inputBoxPass = ui2.NewInputBox(image.Rect(world.ScreenWidth/2-96, 170, world.ScreenWidth/2+96, 202))
	p.btnLogin = ui2.NewButton(image.Rect(world.ScreenWidth/2-32, 240, world.ScreenWidth/2+32, 272), "Login")
	p.btnTestHost = ui2.NewButton(image.Rect(world.ScreenWidth/2+100, 100-32, world.ScreenWidth/2+132, 100), "test")
	p.btnDevHost = ui2.NewButton(image.Rect(world.ScreenWidth/2+100, 100, world.ScreenWidth/2+132, 132), "dev")
	p.btnLocalhost = ui2.NewButton(image.Rect(world.ScreenWidth/2+132, 100, world.ScreenWidth/2+196, 132), "localhost")
	p.cbNginx = ui2.NewCheckBox(world.ScreenWidth/2-32, 206, "connect nginx route")
	p.cbUseLogin = ui2.NewCheckBox(world.ScreenWidth/2+100, 132, "使用dev登陆\n(本地服使用dev库时使用)")

	p.AddChild(p.inputBoxHost)
	p.AddChild(p.inputBoxUser)
	p.AddChild(p.inputBoxPass)
	p.AddChild(p.cbNginx)
	p.AddChild(p.btnLogin)
	p.AddChild(p.btnTestHost)
	p.AddChild(p.btnDevHost)
	p.AddChild(p.btnLocalhost)
	p.AddChild(p.cbUseLogin)

	p.inputBoxHost.MaxChars = 64
	p.inputBoxUser.MaxChars = 64
	p.inputBoxPass.MaxChars = 64
	p.inputBoxHost.Text = host
	p.inputBoxUser.Text = user
	p.inputBoxPass.Text = pass
	p.inputBoxHost.DefaultText = "input server host"
	p.inputBoxUser.DefaultText = "input mail account"
	p.inputBoxPass.DefaultText = "input password"
	p.inputBoxPass.PasswordChar = "*"

	p.inputBoxHost.SetOnPressEnter(func(i *ui2.InputBox) {
		p.inputBoxUser.SetFocused(true)
	})
	p.inputBoxUser.SetOnPressEnter(func(i *ui2.InputBox) {
		p.inputBoxPass.SetFocused(true)
	})
	p.inputBoxPass.SetOnPressEnter(func(i *ui2.InputBox) {
		if p.inputBoxHost.Text == "" {
			p.inputBoxHost.SetFocused(true)
		} else {
			p.btnLogin.Click()
		}
	})
	//p.cbTypeMingfa.SetOnCheckChanged(func(c *CheckBox) {
	//})
	//p.cbUseLogin.SetOnCheckChanged(func(c *CheckBox) {
	//})
	p.btnLogin.SetOnClick(func(b *ui2.Button) {
		host = p.inputBoxHost.Text
		user = p.inputBoxUser.Text
		pass = p.inputBoxPass.Text

		api := "http://" + host
		nginx := p.cbNginx.Checked()
		useDevAPI := p.cbUseLogin.Checked()
		if nginx { // 走nginx路由
			//TODO server用win版的
			// 1.pff_backend\goserver\devops\install\docker\docker-compose.yml 去掉所有networks
			// 2.pff_backend\goserver\devops\install\docker\nginx\nginx.pff.conf 中
			// upstream mc_gateserver {
			//	server gate_server:7671; //改为自己机器对外ip 172.22.*.*:7671;
			//	keepalive 64;
			//}
			//TODO upstream mc_apiserver {
			//	server api_server:7670; //改为自己机器对外ip 172.22.*.*:7671;
			//	keepalive 64;
			//}
		} else { // TODO 直连gate,api
			api = "http://" + host + ":7670"
			host = host + ":7671"
		}
		if useDevAPI {
			api = "http://" + DevHost
		}
		author, err := util.AutoLogin(api, user, pass, gameName)
		if err != nil {
			fmt.Println("AutoLogin.err:", err)
			return
		}
		World.Enter(host, author)
		ebiten.SetWindowTitle(fmt.Sprintf("Home(%s)U(%s)", host, user))
	})

	p.btnDevHost.SetOnClick(func(b *ui2.Button) {
		p.inputBoxHost.Text = DevHost
		p.cbNginx.SetChecked(true)
	})
	p.btnTestHost.SetOnClick(func(b *ui2.Button) {
		p.inputBoxHost.Text = TestHost
		p.cbNginx.SetChecked(true)
	})
	p.btnLocalhost.SetOnClick(func(b *ui2.Button) {
		p.inputBoxHost.Text = LocalHost
		p.cbNginx.SetChecked(false)
	})

	uiLogin = p
	return p
}

func (p *UILogin) OnClose() {
	uiLogin = nil
}
