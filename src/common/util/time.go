package util

import (
	"common/tlog"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func LoadTimeLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if loc == nil {
		tlog.Error(err)
		return time.UTC
	}
	return loc
}

func ParseTime(s string) (time.Time, error) {
	if strings.Index(s, ",") >= 0 {
		return time.Parse(time.RFC1123, s)
	} else if strings.Index(s, "T") >= 0 {
		if len(s) == 19 {
			return time.Time{}, errors.New("inconsistent with the rfc3339")
		}
		return time.Parse(time.RFC3339, s)
	} else if strings.Index(s, ":") >= 0 {
		return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	} else {
		return time.ParseInLocation("2006-01-02", s, time.Local)
	}
}

func FormatBirthday(s string, isUSA bool) string {
	if s == "" {
		return s
	}

	t, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		t, err = time.ParseInLocation("2006/01/02", s, time.Local)
		if err != nil {
			if isUSA {
				t, err = time.ParseInLocation("01/02/2006", s, time.Local)
				if err != nil {
					t, err = time.ParseInLocation("1/2/2006", s, time.Local)
					if err != nil {
						t, err = time.ParseInLocation("02/01/2006", s, time.Local)
						if err != nil {
							t, err = time.ParseInLocation("2/1/2006", s, time.Local)
						}
					}
				}
			} else {
				t, err = time.ParseInLocation("02/01/2006", s, time.Local)
				if err != nil {
					t, err = time.ParseInLocation("2/1/2006", s, time.Local)
					if err != nil {
						t, err = time.ParseInLocation("01/02/2006", s, time.Local)
						if err != nil {
							t, err = time.ParseInLocation("1/2/2006", s, time.Local)
						}
					}
				}
			}
		}
	}
	if err == nil {
		y, m, d := t.Date()
		return fmt.Sprintf("%4d-%02d-%02d", y, m, d)
	} else {
		return ""
	}
}

func FormatShortDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%4d%02d%02d", y, m, d)
}

func FormatDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%4d-%02d-%02d", y, m, d)
}

func FormatTime(t time.Time) string {
	y, m, d := t.Date()
	hour, minute, second := t.Clock()
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", y, m, d, hour, minute, second)
}

func GetZeroUnixTime(unixTime int64, locName string) (int64, int64) {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, dd := t.Date()
	tt := time.Date(yy, mm, dd, 0, 0, 0, 0, loc)
	return tt.Unix(), unixTime - tt.Unix()
}

func GetMidnightUnixTime(unixTime int64, locName string) (int64, int64) {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)
	yy, mm, dd := t.Date()
	tt := time.Date(yy, mm, dd+1, 0, 0, 0, 0, loc)
	return tt.Unix(), tt.Unix() - unixTime
}

func GetNextMidnightUnixTime(unixTime int64, locName string) (int64, int64) {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)
	yy, mm, dd := t.Date()
	tt := time.Date(yy, mm, dd+2, 0, 0, 0, 0, loc)
	return tt.Unix(), tt.Unix() - unixTime
}

func IsSameDay(t1, t2 int64, locName string) bool {
	loc := LoadTimeLocation(locName)
	tt1 := time.Unix(t1, 0).In(loc)
	tt2 := time.Unix(t2, 0).In(loc)

	y1, m1, d1 := tt1.Date()
	y2, m2, d2 := tt2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func IsSameMonth(t1, t2 int64, locName string) bool {
	loc := LoadTimeLocation(locName)
	tt1 := time.Unix(t1, 0).In(loc)
	tt2 := time.Unix(t2, 0).In(loc)

	y1, m1, _ := tt1.Date()
	y2, m2, _ := tt2.Date()
	return y1 == y2 && m1 == m2
}

func IsSameWeek(t1, t2 int64, locName string) bool {
	loc := LoadTimeLocation(locName)
	tt1 := time.Unix(t1, 0).In(loc)
	tt2 := time.Unix(t2, 0).In(loc)

	y1, w1 := tt1.ISOWeek()
	y2, w2 := tt2.ISOWeek()
	return y1 == y2 && w1 == w2
}

// 求相差天数
func GetDiffDays(t1, t2 int64, locName string) int64 {
	loc := LoadTimeLocation(locName)
	tt1 := time.Unix(t1, 0).In(loc)
	tt2 := time.Unix(t2, 0).In(loc)

	y1, m1, d1 := tt1.Date()
	y2, m2, d2 := tt2.Date()
	var dd1, dd2 time.Time

	dd1 = time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC)
	dd2 = time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC)

	return int64(dd1.Sub(dd2).Hours() / 24)
}

// 获取星期几
func ConvertWeek(t int64, locName string) int {
	loc := LoadTimeLocation(locName)
	t1 := time.Unix(t, 0).In(loc)

	week := t1.Weekday()
	return int(week)
}

// 获取下周（一）开始的 时间戳
func NextFirstDayOfWeek(t int64, locName string) (endTime int64) {
	loc := LoadTimeLocation(locName)
	t1 := time.Unix(t, 0).In(loc)

	week := t1.Weekday()
	offset := int(time.Monday - week)
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	year, month, day := t1.Date()

	thisWeek := time.Date(year, month, day, 0, 0, 0, 0, loc)
	endTime = thisWeek.AddDate(0, 0, offset+7).Unix()
	return
}

// 获取下月 1 号开始的时间戳
func NextFirstDayOfMonth(t int64, locName string) (endTime int64) {
	loc := LoadTimeLocation(locName)
	t1 := time.Unix(t, 0).In(loc)

	year, month, _ := t1.Date()

	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	endTime = thisMonth.AddDate(0, 1, 0).Unix()
	return
}

// 获取年月日
func GetYmd(t int64, locName string) (int, int, int) {
	loc := LoadTimeLocation(locName)
	t1 := time.Unix(t, 0).In(loc)

	year, month, day := t1.Date()
	return year, int(month), day
}

// 获取小时
func GetHour(t int64, locName string) int {
	loc := LoadTimeLocation(locName)
	t1 := time.Unix(t, 0).In(loc)

	return t1.Hour()
}

func GetDayBeginTimeWithOffset(unixTime int64, offsetHours int, locName string) time.Time {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, dd := t.Date()
	beginTime := time.Date(yy, mm, dd, offsetHours, 0, 0, 0, loc)

	if t.Unix() < beginTime.Unix() {
		beginTime = beginTime.AddDate(0, 0, -1)
	}
	return beginTime
}

func GetWeekBeginTimeWithOffset(unixTime int64, offsetWeekday time.Weekday, offsetHours int, locName string) time.Time {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, dd := t.Date()
	beginTime := time.Date(yy, mm, dd, offsetHours, 0, 0, 0, loc)

	weekday := int(beginTime.Weekday())
	days := 0
	if weekday == 0 {
		days = 6
	} else {
		days = weekday - 1
	}

	offsetDays := 0
	if offsetWeekday == time.Sunday {
		offsetDays = 6
	} else {
		offsetDays = int(offsetWeekday) - 1
	}

	beginTime = beginTime.AddDate(0, 0, -days+offsetDays)
	if t.Unix() < beginTime.Unix() {
		beginTime = beginTime.AddDate(0, 0, -7)
	}

	return beginTime
}

func GetMonthBeginTimeWithOffset(unixTime int64, offsetDays, offsetHours int, locName string) time.Time {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, _ := t.Date()
	beginTime := time.Date(yy, mm, offsetDays, offsetHours, 0, 0, 0, loc)
	if t.Unix() < beginTime.Unix() {
		beginTime = beginTime.AddDate(0, -1, 0)
	}
	return beginTime
}

func FindNearestDailyTimePoint(unixTime int64, points []int32, locName string) int64 {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, dd := t.Date()
	zero := time.Date(yy, mm, dd, 0, 0, 0, 0, loc)

	var nearestTime *time.Time
	for _, p := range points {
		t1 := zero.Add(time.Second * time.Duration(p))
		if t1.Before(t) {
			t1 = t1.AddDate(0, 0, 1)
		}
		if nearestTime == nil || t1.Before(*nearestTime) {
			nearestTime = &t1
		}
	}

	//fmt.Println(*nearestTime)
	return nearestTime.Unix()
}

func FindNearestWeeklyTimePoint[T any](unixTime int64, weekdays []int32, points []T, fn func(p T) int32, locName string) (int64, int) {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	weekday := int(t.Weekday())
	days := 0
	if weekday == 0 {
		days = 6
	} else {
		days = weekday - 1
	}
	monday := t.AddDate(0, 0, -days)
	yy, mm, dd := monday.Date()
	mondayZero := time.Date(yy, mm, dd, 0, 0, 0, 0, loc)

	var idx int
	var nearestTime *time.Time
	for _, wd := range weekdays {
		if wd <= 0 || wd > 7 {
			wd = 7
		}
		weekdayZero := mondayZero.AddDate(0, 0, int(wd)-1)

		for i, point := range points {
			p := fn(point)
			t1 := weekdayZero.Add(time.Second * time.Duration(p))
			if !t1.After(t) {
				t1 = t1.AddDate(0, 0, 7)
			}
			if nearestTime == nil || t1.Before(*nearestTime) {
				idx = i
				nearestTime = &t1
			}
		}
	}

	//fmt.Println(*nearestTime)
	//fmt.Println(idx)
	return nearestTime.Unix(), idx
}

// FindNearestWeeklyRandInterval finds the nearest random interval in seconds between secsMin and secsMax on the specified weekdays.
// The interval is calculated from the given unixTime and offsetInit.
// locName is the name of the location to use for time zone information.
func FindNearestWeeklyRandInterval(unixTime, offsetInit int64, weekdays []int32, secsMin, secsMax int32, locName string) int64 {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	weekdayMap := make(map[int32]bool)
	for _, wd := range weekdays {
		if wd == 0 {
			wd = 7
		}
		weekdayMap[wd] = true
	}

	for i := 0; i <= 7; i++ { // search for the next 7 days
		t1 := t.AddDate(0, 0, i)
		weekday := int32(t1.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		if weekdayMap[weekday] == false {
			continue
		}

		t1Zero, _ := GetZeroUnixTime(t1.Unix(), locName)
		offset := int64(rand.Int31n(secsMax-secsMin+1) + secsMin)
		if i == 0 {
			t1Midnight, _ := GetMidnightUnixTime(t1.Unix(), locName)
			for t1Zero+offsetInit+offset < t1Midnight {
				if t1Zero+offsetInit+offset > unixTime {
					return t1Zero + offsetInit + offset
				}
				offset += int64(rand.Int31n(secsMax-secsMin+1) + secsMin)
			}
			continue
		}

		return t1Zero + offset
	}

	return 0
}

func FindNearestMonthlyTimePoint[T any](unixTime int64, days []int32, points []T, fn func(p T) int32, locName string) (int64, int) {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	yy, mm, _ := t.Date()

	var idx int
	var nearestTime *time.Time
	for _, day := range days {
		dayZero := time.Date(yy, mm, int(day), 0, 0, 0, 0, loc)
		for i, point := range points {
			p := fn(point)
			t1 := dayZero.Add(time.Second * time.Duration(p))
			if t1.Before(t) {
				t1 = t1.AddDate(0, 1, 0)
			}
			if nearestTime == nil || t1.Before(*nearestTime) {
				idx = i
				nearestTime = &t1
			}
		}
	}

	//fmt.Println(*nearestTime)
	//fmt.Println(idx)
	return nearestTime.Unix(), idx
}

func IsInDailyTimeRange[T any](unixTime int64, ranges []T, fn func(r T) (int32, int32), locName string) bool {
	zero, _ := GetZeroUnixTime(unixTime, locName)
	for _, r := range ranges {
		begin, end := fn(r)
		if zero+int64(begin) <= unixTime && unixTime < zero+int64(end) {
			return true
		}
	}

	return false
}

func IsInWeeklyTimeRange[T any](unixTime int64, weekdays []int32, ranges []T, fn func(r T) (int32, int32), locName string) bool {
	loc := LoadTimeLocation(locName)
	t := time.Unix(unixTime, 0).In(loc)

	weekday := int32(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	var find = false
	for _, wd := range weekdays {
		if wd == 0 {
			wd = 7
		}
		if wd == weekday {
			find = true
			break
		}
	}
	if !find {
		return false
	}

	zero, _ := GetZeroUnixTime(unixTime, locName)
	for _, r := range ranges {
		begin, end := fn(r)
		if zero+int64(begin) <= unixTime && unixTime < zero+int64(end) {
			return true
		}
	}

	return false
}
