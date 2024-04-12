package logic

import (
	"client0/logic/asset"
	"client0/logic/ebiten/ui"
	"common/qimen"
	"errors"
	"fmt"
	"github.com/6tail/lunar-go/SolarUtil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"strconv"
	"time"
)

// CelestialStem

func init() {
	f, _ := asset.LoadFont("font/lana_pixel.ttf", &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	ui.SetDefaultUIFont(f)
}

type UIQiMen struct {
	ui.BaseUI
	inputSYear  *ui.InputBox
	inputSMonth *ui.InputBox
	inputSDay   *ui.InputBox
	inputSHour  *ui.InputBox
	inputSMin   *ui.InputBox

	inputLYear   *ui.InputBox
	inputLMonth  *ui.InputBox
	inputLDay    *ui.InputBox
	inputLHour   *ui.InputBox
	inputYearRB  *ui.InputBox
	inputMonthRB *ui.InputBox
	inputDayRB   *ui.InputBox
	inputHourRB  *ui.InputBox
	inputJieQi   *ui.InputBox

	cbNginx    *ui.CheckBox
	cbUseLogin *ui.CheckBox

	btnCalc      *ui.Button
	btnPreHour2  *ui.Button
	btnNextHour2 *ui.Button

	textGong []*ui.TextBox
}

var uiQiMen *UIQiMen

func UIShowQiMen(width, height int) {
	if uiQiMen == nil {
		uiQiMen = NewUIQiMen(width, height)
		ui.ActiveUI(uiQiMen)
	}
}
func UIHideQiMen() {
	if uiQiMen != nil {
		ui.CloseUI(uiQiMen)
		uiQiMen = nil
	}
}

func NewUIQiMen(width, height int) *UIQiMen {
	//cx, cy := width/2, height/2 //win center
	p := &UIQiMen{}
	p.inputSYear = ui.NewInputBox(image.Rect(32, 32, 32+64, 64))
	p.inputSMonth = ui.NewInputBox(image.Rect(32+72, 32, 32+72+64, 64))
	p.inputSDay = ui.NewInputBox(image.Rect(32+72*2, 32, 32+72*2+64, 64))
	p.inputSHour = ui.NewInputBox(image.Rect(32+72*3, 32, 32+72*3+64, 64))
	p.inputSMin = ui.NewInputBox(image.Rect(32+72*4, 32, 32+72*4+64, 64))
	p.btnCalc = ui.NewButton(image.Rect(32+72*5, 32, 32+72*5+64, 64), "Go")

	p.inputLYear = ui.NewInputBox(image.Rect(32, 32+32, 32+64, 32+64))
	p.inputLMonth = ui.NewInputBox(image.Rect(32+72, 32+32, 32+72+64, 32+64))
	p.inputLDay = ui.NewInputBox(image.Rect(32+72*2, 32+32, 32+72*2+64, 32+64))
	p.inputLHour = ui.NewInputBox(image.Rect(32+72*3, 32+32, 32+72*3+64, 32+64))
	p.btnPreHour2 = ui.NewButton(image.Rect(32+72*4, 32+32, 32+72*4+64, 32+64), "Pre")
	p.btnNextHour2 = ui.NewButton(image.Rect(32+72*5, 32+32, 32+72*5+64, 32+64), "Next")

	p.inputYearRB = ui.NewInputBox(image.Rect(32, 32+32*2, 32+64, 32*2+64))
	p.inputMonthRB = ui.NewInputBox(image.Rect(32+72, 32+32*2, 32+72+64, 32*2+64))
	p.inputDayRB = ui.NewInputBox(image.Rect(32+72*2, 32+32*2, 32+72*2+64, 32*2+64))
	p.inputHourRB = ui.NewInputBox(image.Rect(32+72*3, 32+32*2, 32+72*3+64, 32*2+64))
	p.inputJieQi = ui.NewInputBox(image.Rect(32+72*4, 32+32*2, 32+72*4+128+8, 32*2+64))
	//p.cbNginx = ui.NewCheckBox(cx-32, 206, "XXX")

	gongOffset := [][]int{{0, 0},
		{150, 300}, {300, 0}, {0, 150},
		{0, 0}, {150, 150}, {300, 300},
		{300, 150}, {0, 300}, {150, 0},
	}
	p.textGong = append(p.textGong, nil)
	for i := 1; i <= 9; i++ {
		offX, offZ := gongOffset[i][0], gongOffset[i][1]
		txtGong := ui.NewTextBox(image.Rect(32+offX, 64+32*2+offZ, 32+150+offX, 64+32*2+150+offZ))
		txtGong.SetText(i)
		p.textGong = append(p.textGong, txtGong)
		p.AddChild(txtGong)
	}

	p.AddChild(p.inputSYear)
	p.AddChild(p.inputSMonth)
	p.AddChild(p.inputSDay)
	p.AddChild(p.inputSHour)
	p.AddChild(p.inputSMin)
	p.AddChild(p.btnCalc)

	p.AddChild(p.inputLYear)
	p.AddChild(p.inputLMonth)
	p.AddChild(p.inputLDay)
	p.AddChild(p.inputLHour)
	p.AddChild(p.btnPreHour2)
	p.AddChild(p.btnNextHour2)

	p.AddChild(p.inputYearRB)
	p.AddChild(p.inputMonthRB)
	p.AddChild(p.inputDayRB)
	p.AddChild(p.inputHourRB)
	p.AddChild(p.inputJieQi)

	//p.AddChild(p.cbNginx)
	//p.AddChild(p.btnTestHost)
	//p.AddChild(p.btnDevHost)
	//p.AddChild(p.btnLocalhost)

	p.inputSYear.MaxChars = 4
	p.inputSMonth.MaxChars = 2
	p.inputSDay.MaxChars = 2
	p.inputSHour.MaxChars = 2
	p.inputSMin.MaxChars = 2

	p.inputSYear.DefaultText = "year"
	p.inputSMonth.DefaultText = "month"
	p.inputSDay.DefaultText = "day"
	p.inputSHour.DefaultText = "hour"
	p.inputSMin.DefaultText = "minute"

	p.btnCalc.SetOnClick(func(b *ui.Button) {
		year, _ := strconv.Atoi(p.inputSYear.Text)
		month, _ := strconv.Atoi(p.inputSMonth.Text)
		day, _ := strconv.Atoi(p.inputSDay.Text)
		hour, _ := strconv.Atoi(p.inputSHour.Text)
		minute, _ := strconv.Atoi(p.inputSMin.Text)
		p.Apply(year, month, day, hour, minute)
	})
	p.btnPreHour2.SetOnClick(func(b *ui.Button) {
		year, _ := strconv.Atoi(p.inputSYear.Text)
		month, _ := strconv.Atoi(p.inputSMonth.Text)
		day, _ := strconv.Atoi(p.inputSDay.Text)
		hour, _ := strconv.Atoi(p.inputSHour.Text)
		minute, _ := strconv.Atoi(p.inputSMin.Text)
		hour -= 2
		if hour < 0 {
			hour += 24
			day--
			if day == 0 {
				month--
				if month == 0 {
					month = 12
					year--
				}
				day = SolarUtil.GetDaysOfMonth(year, month)
			}
			if 1582 == year && 10 == month && day == 14 {
				day = 4
			}
		}
		p.Apply(year, month, day, hour, minute)
	})
	p.btnNextHour2.SetOnClick(func(b *ui.Button) {
		year, _ := strconv.Atoi(p.inputSYear.Text)
		month, _ := strconv.Atoi(p.inputSMonth.Text)
		day, _ := strconv.Atoi(p.inputSDay.Text)
		hour, _ := strconv.Atoi(p.inputSHour.Text)
		minute, _ := strconv.Atoi(p.inputSMin.Text)
		hour += 2
		if hour > 23 {
			hour -= 24
			day++
			if day > SolarUtil.GetDaysOfMonth(year, month) {
				day = 1
				month++
				if month > 12 {
					month = 1
					year++
				}
			}
			if 1582 == year && 10 == month && day == 5 {
				day = 15
			}
		}
		p.Apply(year, month, day, hour, minute)
	})

	p.inputLYear.Editable = false
	p.inputLMonth.Editable = false
	p.inputLDay.Editable = false
	p.inputLHour.Editable = false

	p.inputYearRB.Editable = false
	p.inputMonthRB.Editable = false
	p.inputDayRB.Editable = false
	p.inputHourRB.Editable = false
	p.inputJieQi.Editable = false

	t := time.Now()
	p.Apply(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())

	uiQiMen = p
	return p
}

func (p *UIQiMen) OnClose() {
	uiQiMen = nil
}

func (p *UIQiMen) checkDate(year, month, day, hour, minute int) error {
	if month < 1 || month > 12 {
		return errors.New(fmt.Sprintf("wrong month %v", month))
	}
	if day < 1 || day > 31 {
		return errors.New(fmt.Sprintf("wrong day %v", day))
	}
	if 1582 == year && 10 == month {
		if day > 4 && day < 15 {
			return errors.New(fmt.Sprintf("wrong solar year %v month %v day %v", year, month, day))
		}
	} else {
		if day > SolarUtil.GetDaysOfMonth(year, month) {
			return errors.New(fmt.Sprintf("wrong solar year %v month %v day %v", year, month, day))
		}
	}
	if hour < 0 || hour > 23 {
		return errors.New(fmt.Sprintf("wrong hour %v", hour))
	}
	if minute < 0 || minute > 59 {
		return errors.New(fmt.Sprintf("wrong minute %v", minute))
	}
	return nil
}
func (p *UIQiMen) Apply(year, month, day, hour, minute int) {
	pan, err := qimen.NewPan(year, month, day, hour, minute)
	if err != nil {
		UIShowMsgBox("时间不对", "确定", "取消", func(b *ui.Button) {
		}, func(b *ui.Button) {})
	}
	//pan.DayArr
	p.inputSYear.SetText(pan.SolarYear)
	p.inputSMonth.SetText(pan.SolarMonth)
	p.inputSDay.SetText(pan.SolarDay)
	p.inputSHour.SetText(pan.SolarHour)
	p.inputSMin.SetText(pan.SolarMinute)

	p.inputLYear.SetText(pan.LunarYearC)
	p.inputLMonth.SetText(pan.LunarMonthC)
	p.inputLDay.SetText(pan.LunarDayC)
	p.inputLHour.SetText(pan.LunarHourC)

	p.inputYearRB.SetText(pan.YearRB)
	p.inputMonthRB.SetText(pan.MonthRB)
	p.inputDayRB.SetText(pan.DayRB)
	p.inputHourRB.SetText(pan.HourRB)

	p.inputJieQi.Text = pan.JuText

	for i := 1; i <= 9; i++ {

		p.textGong[i].Text = fmt.Sprintf("神\n星\n门\n地\n\n%s%d", qimen.Gua8Gong9[i], i)
	}
}
