package qimen

import (
	"errors"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/SolarUtil"
	"github.com/6tail/lunar-go/calendar"
)

// 奇门遁甲宫格
type QiMenGong struct {
	Diagram8 int //八卦(固定)
	EHs36    int //地盘
	HHs36    int //天盘
	EStar    int //地九星(固定)
	HStar    int //天九星1~9
	Door8    int //八门
	God9     int //九神1~9
}

type QiMenPan struct {
	//SolarTime time.Time
	SolarYear   int
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

	//Y_HStems    int //年天干1~10
	//Y_EBranches int //年地支1~12
	//M_HStems    int //月天干1~10
	//M_EBranches int //月地支1~12
	//D_HStems    int //日天干1~10
	//D_EBranches int //日地支1~12
	//H_HStems    int //时天干1~10
	//H_EBranches int //时地支1~12
	//Constellation28 int            //星宿1～28

	YearRB  string //年干支
	MonthRB string //月干支
	DayRB   string //日干支
	HourRB  string //时干支
	JuText  string //局文本
	JieQi   int    //节气1~24
	Yuan3   int    //三元1~3
	Ju      int    //格局-1~-9,1~9

	Duty        int //值序
	DutyGod     int //值符
	DutyGodPos  int //值符落宫
	DutyDoor    int //值使
	DutyDoorPos int //值使落宫

	FlyArr   [10]QiMenGong //九宫飞盘格
	RollArr  [10]QiMenGong //九宫转盘格
	RollFly  [10]QiMenGong //半飞半转盘
	DayArr   [10]QiMenGong //日家奇门盘
	MonthArr [10]QiMenGong //月家奇门盘
	YearArr  [10]QiMenGong //年家奇门盘

}

func getQiMenYunIndex(dayGanZhi string) int {
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
	jqi := jieQiIndex[jieQi]
	return qiMenJu[jqi][yuan3Idx-1]
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
	yuanIdx := getQiMenYunIndex(c8[2])
	yuanName := sanYuanName[yuanIdx]
	juIdx := getQiMenJuIndex(cal.GetPrevJieQi().GetName(), yuanIdx)
	var juName string
	if juIdx < 0 {
		juName = fmt.Sprintf("阴%d局", juIdx)
	} else {
		juName = fmt.Sprintf("阳%d局", juIdx)
	}
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
		Ju:          juIdx,
		Yuan3:       yuanIdx,

		JuText: fmt.Sprintf("%s %s %s", cal.GetPrevJieQi().GetName(), yuanName, juName),
	}
	return &p, nil
}
