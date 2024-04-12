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

	textLYear   *ui.InputBox
	textLMonth  *ui.InputBox
	textLDay    *ui.InputBox
	textLHour   *ui.InputBox
	textYearRB  *ui.InputBox
	textMonthRB *ui.InputBox
	textDayRB   *ui.InputBox
	textHourRB  *ui.InputBox
	textJu      *ui.InputBox

	cbType *ui.CheckBox

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

	p.textLYear = ui.NewInputBox(image.Rect(32, 32+32, 32+64, 32+64))
	p.textLMonth = ui.NewInputBox(image.Rect(32+72, 32+32, 32+72+64, 32+64))
	p.textLDay = ui.NewInputBox(image.Rect(32+72*2, 32+32, 32+72*2+64, 32+64))
	p.textLHour = ui.NewInputBox(image.Rect(32+72*3, 32+32, 32+72*3+64, 32+64))
	p.btnPreHour2 = ui.NewButton(image.Rect(32+72*4, 32+32, 32+72*4+64, 32+64), "Pre")
	p.btnNextHour2 = ui.NewButton(image.Rect(32+72*5, 32+32, 32+72*5+64, 32+64), "Next")

	p.textYearRB = ui.NewInputBox(image.Rect(32, 32+32*2, 32+64, 32*2+64))
	p.textMonthRB = ui.NewInputBox(image.Rect(32+72, 32+32*2, 32+72+64, 32*2+64))
	p.textDayRB = ui.NewInputBox(image.Rect(32+72*2, 32+32*2, 32+72*2+64, 32*2+64))
	p.textHourRB = ui.NewInputBox(image.Rect(32+72*3, 32+32*2, 32+72*3+64, 32*2+64))
	p.textJu = ui.NewInputBox(image.Rect(32+72*4, 32+32*2, 32+72*4+128+8+256, 32*2+64))
	//p.cbType = ui.NewCheckBox(cx-32, 206, "XXX")

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

	p.AddChild(p.textLYear)
	p.AddChild(p.textLMonth)
	p.AddChild(p.textLDay)
	p.AddChild(p.textLHour)
	p.AddChild(p.btnPreHour2)
	p.AddChild(p.btnNextHour2)

	p.AddChild(p.textYearRB)
	p.AddChild(p.textMonthRB)
	p.AddChild(p.textDayRB)
	p.AddChild(p.textHourRB)
	p.AddChild(p.textJu)

	//p.AddChild(p.cbType)
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

	p.textLYear.Editable = false
	p.textLMonth.Editable = false
	p.textLDay.Editable = false
	p.textLHour.Editable = false

	p.textYearRB.Editable = false
	p.textMonthRB.Editable = false
	p.textDayRB.Editable = false
	p.textHourRB.Editable = false
	p.textJu.Editable = false

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

	p.textLYear.SetText(pan.LunarYearC)
	p.textLMonth.SetText(pan.LunarMonthC)
	p.textLDay.SetText(pan.LunarDayC)
	p.textLHour.SetText(pan.LunarHourC)

	p.textYearRB.SetText(pan.YearRB)
	p.textMonthRB.SetText(pan.MonthRB)
	p.textDayRB.SetText(pan.DayRB)
	p.textHourRB.SetText(pan.HourRB)

	p.textJu.Text = pan.JuText

	for i := 1; i <= 9; i++ {
		p.textGong[i].Text = pan.Gongs[i].FmtText
	}
}
