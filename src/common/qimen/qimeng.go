package qimen

import (
	"common/util"
	"errors"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/SolarUtil"
	"github.com/6tail/lunar-go/calendar"
)

// 奇门遁甲宫格
type QiMenGong struct {
	Idx int //洛书宫数

	EHs36 string //地盘奇仪
	HHs36 string //天盘奇仪
	HStar string //天九星1~9
	Door8 string //八门
	God9  string //九神1~9

	FmtText string
}

type QiMenPan struct {
	//SolarTime time.Time

	SolarYear   int //1900-2100
	SolarMonth  int //1-12
	SolarDay    int //1-31
	SolarHour   int //0-23
	SolarMinute int //分

	lunarYear    int //农历年
	lunarMonth   int //农历月 1~12 闰-1~-12
	lunarDay     int //农历日 1~30
	lunarHour    int //农历时
	lunarQuarter int //农历刻
	LunarYearC   string
	LunarMonthC  string
	LunarDayC    string
	LunarHourC   string

	//days       int //1900.1.1起总Days
	//dayPeriod  int //日旬1~6
	//hourPeriod int //时旬1~6

	//Y_HStems    int //年干1~10
	//Y_EBranches int //年支1~12
	//M_HStems    int //月干1~10
	//M_EBranches int //月支1~12
	//D_HStems    int //日干1~10
	//D_EBranches int //日支1~12
	//H_HStems    int //时干1~10
	//H_EBranches int //时支1~12
	//Constellation28 int            //星宿1～28

	YearRB    string //年干支
	MonthRB   string //月干支
	DayRB     string //日干支
	HourRB    string //时干支
	JieQiName string //节气文本
	Yuan3     int    //三元1~3
	Ju        int    //格局-1~-9,1~9

	ShiXun      string //时辰旬首
	Duty        int    //值序
	DutyStar    string //值符
	DutyStarPos int    //值符落宫
	DutyDoor    string //值使
	DutyDoorPos int    //值使落宫

	JuText string        //局文本
	Gongs  [10]QiMenGong //九宫飞盘格
	//FlyArr  [10]QiMenGong //九宫飞盘格
	//RollArr [10]QiMenGong //九宫转盘格
	//RollFly [10]QiMenGong //半飞半转盘

	DayArr   [10]QiMenGong //日家奇门盘
	MonthArr [10]QiMenGong //月家奇门盘
	YearArr  [10]QiMenGong //年家奇门盘

}

func (p *QiMenPan) calcGong() {
	g9 := &p.Gongs
	for i := 1; i <= 9; i++ {
		g9[i].Idx = i
		//g9[i].Diagram8 = _Gua8In9[i]
		//g9[i].HomeStar = QiMenStar9[i]
		//g9[i].HomeDoor = QiMenDoor9[i]
	}
	//地盘
	gongStart := util.Abs(p.Ju)
	if p.Ju > 0 { //阳遁顺仪奇逆布
		for i := gongStart; i <= 9; i++ {
			g9[i].EHs36 = _Qm3Q6Y[i-gongStart]
		}
		for i := 1; i < gongStart; i++ {
			g9[i].EHs36 = _Qm3Q6Y[i-gongStart+9]
		}
	} else { //阴遁逆仪奇顺行
		for i := gongStart; i > 0; i-- {
			g9[i].EHs36 = _Qm3Q6Y[gongStart-i]
		}
		for i := 1; i < gongStart; i-- {
			g9[i].EHs36 = _Qm3Q6Y[i-gongStart-9]
		}
	}
	//值符
	for i := 1; i <= 9; i++ {
		if g9[i].EHs36 == _HideJia[p.ShiXun] {
			p.Duty = i
			p.DutyStar = _QiMenStar9[i]
			//p.DutyStarPos int    //值符落宫
			p.DutyDoor = _QiMenDoor9[i] // TODO if 转盘值使寄坤宫
			//p.DutyDoorPos int    //值使落宫
			break
		}
	}

	//fmt
	for i := 1; i <= 9; i++ {
		g9[i].FmtText = fmt.Sprintf("\n九神\n\n九星    天盘\n\n八门    %s\n\n%s %s",
			g9[i].EHs36, _Gua8In9[i], LunarUtil.NUMBER[i])
	}
}

func getQiMenYuan3Index(dayGanZhi string) int {
	jiaZiIndex := LunarUtil.GetJiaZiIndex(dayGanZhi)
	qiMenYuanIdx := jiaZiIndex % 15
	if qiMenYuanIdx < 5 {
		return 1
	} else if qiMenYuanIdx < 10 {
		return 2
	}
	return 3
}

func getQiMenJuIndex(jieQi string, yuan3Idx int) int {
	jqi := _JieQiIndex[jieQi]
	return _QiMenJu[jqi][yuan3Idx-1]
}

// 返回solar年的第n(1小寒)个节气进入时间
func getTermTime(year, n int) int64 {
	t := int64(31556925974.7*float64(year-1900)/1000) + int64(termData[n-1]*60-2208549300)
	return t
}

func checkDate(year, month, day, hour, minute int) error {
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

func NewPan(year int, month int, day int, hour int, minute int) (*QiMenPan, error) {
	if err := checkDate(year, month, day, hour, minute); err != nil {
		return nil, err
	}
	cal := calendar.NewLunarFromSolar(calendar.NewSolar(year, month, day, hour, minute, 0))
	c8 := cal.GetBaZi()
	dayGanZhi := c8[2]
	shiGanZhi := c8[3]
	if hour == 23 { //晚子时日柱作次日
		di := LunarUtil.GetJiaZiIndex(dayGanZhi) + 1
		if di > 59 {
			di -= 60
		}
		dayGanZhi = LunarUtil.JIA_ZI[di]
	}
	dayYuanIdx := getQiMenYuan3Index(c8[2])
	yuanName := _Yuan3Name[dayYuanIdx]
	jieQiName := cal.GetPrevJieQi().GetName()
	juIdx := getQiMenJuIndex(jieQiName, dayYuanIdx)
	var juName string
	if juIdx < 0 {
		juName = fmt.Sprintf("阴%d局", juIdx)
	} else {
		juName = fmt.Sprintf("阳%d局", juIdx)
	}
	shiXun := LunarUtil.GetXun(shiGanZhi)
	p := QiMenPan{
		SolarYear:   year,
		SolarMonth:  month,
		SolarDay:    day,
		SolarHour:   hour,
		SolarMinute: minute,
		lunarYear:   cal.GetYear(),
		lunarMonth:  cal.GetMonth(),
		lunarDay:    cal.GetYear(),
		lunarHour:   cal.GetHour(),
		LunarYearC:  cal.GetYearInChinese(),
		LunarMonthC: cal.GetMonthInChinese() + "月",
		LunarDayC:   cal.GetDayInChinese(),
		LunarHourC:  shiGanZhi[len(shiGanZhi)/2:] + "时",
		YearRB:      c8[0],
		MonthRB:     c8[1],
		DayRB:       dayGanZhi,
		HourRB:      shiGanZhi,
		JieQiName:   jieQiName,
		Ju:          juIdx,
		Yuan3:       dayYuanIdx,
		ShiXun:      shiXun,
	}
	p.calcGong()
	p.JuText = fmt.Sprintf("%s%s %s %s遁%s 值符%s落 值使%s落", jieQiName, yuanName, juName,
		shiXun, _HideJia[shiXun], p.DutyStar, p.DutyDoor)

	return &p, nil
}
