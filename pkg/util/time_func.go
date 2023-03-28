package util

import "time"

const (
	FormatDate      = "2006-01-02 15:04:05"
	FormatSecondStr = "20060102150405"
	FormatDay       = "2006-01-02"
)

const (
	OneMonthDay = 31
	DefaultTime = "1980-01-01 00:00:01"
)

func StringToTime(time_ string) time.Time {
	res, err := time.ParseInLocation(FormatDate, time_, time.Local)
	if err != nil {
		return time.Time{} // 生成个默认的，降级
	}
	return res
}

func StringDayToTime(time_ string) time.Time {
	res, err := time.ParseInLocation(FormatDay, time_, time.Local)
	if err != nil {
		return time.Time{} // 生成个默认的，降级
	}
	return res
}

func IsDefaultTime(t time.Time) bool {
	return t == GetDefaultTime()
}

func GetDefaultTime() time.Time {
	location, _ := time.LoadLocation("PRC")
	defaultTime, _ := time.ParseInLocation(FormatDate, DefaultTime, location)

	return defaultTime
}

func GetUnixtimeDiff(a, b time.Time) int64 {

	if a.Unix() > b.Unix() {
		return a.Unix() - b.Unix()
	}

	return b.Unix() - a.Unix()
}

func ZeroTime() time.Time {
	return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
}

func NilTime() time.Time {
	// 目前系统未设置时区
	location, _ := time.LoadLocation("PRC")
	nilTime, _ := time.ParseInLocation(FormatDate, "0001-01-01 00:00:00", location)
	return nilTime
}

func IsSameDay(date1 time.Time, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func GetDurationDays(startTime int64, endTime int64) []interface{} {
	var result []interface{}
	dayTime := int64(3600 * 24)
	nowTime := time.Now().Unix()
	if endTime == 0 || endTime > nowTime {
		endTime = nowTime
	}
	if startTime == 0 || startTime > nowTime {
		startTime = nowTime
	}
	for {
		if startTime > endTime {
			break
		}
		result = append(result, time.Unix(startTime, 0).Format("2006-01-02"))
		startTime += dayTime
	}
	return result
}

func FormatTime(timestamp int64) string {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func GetCurrentMonthDays() []interface{} {
	var result []interface{}
	monthFirstDay := time.Now().Format("2006-01") + "-01"
	startTime, err := time.Parse("2006-01-02", monthFirstDay)
	if err == nil {
		month := startTime.Month()
		day := time.Now().Day()
		for {
			if startTime.Month() != month || startTime.Day() > day {
				break
			}
			result = append(result, startTime.Format("2006-01-02"))
			startTime = startTime.AddDate(0, 0, 1)
		}
	}
	return result
}

func GetCurrentMonthDaysByPage(pageIndex int, pageSize int) []interface{} {
	var result []interface{}
	n := 0 - (pageIndex * pageSize) + 1
	startTime := time.Now()
	month := startTime.Month()
	day := startTime.Day()
	startTime = startTime.AddDate(0, 0, n)
	i := 0
	for {
		if i >= pageSize {
			break
		}
		if startTime.Month() == month {
			result = append(result, startTime.Format("2006-01-02"))
			if startTime.Day() > day {
				break
			}
		}
		startTime = startTime.AddDate(0, 0, 1)
		i++
	}
	return result
}

/*
时间段 转 日期
*/
func TimeDurationToDateSlice(s, e time.Time, layout string) []string {
	var dateS []string
	if e.Before(s) {
		ee := e
		e = s
		s = ee
	}
	for s.Before(e) {
		dateS = append(dateS, s.Format(layout))
		s = s.AddDate(0, 0, 1)
	}
	return dateS
}

func TimeToString(t time.Time, format string) (timeString string) {
	if !NilTime().Equal(t) {
		timeString = t.Format(format)
	}
	return timeString
}

func NilTimeLocal() time.Time {
	nilTime, _ := time.ParseInLocation(FormatDate, "0001-01-01 00:00:00", time.Local)
	return nilTime
}

func TimeSec2Time(sec int64) time.Time {
	if sec == 0 {
		return NilTimeLocal()
	}
	return time.Unix(sec, 0)
}

func DayStartAndEndTime(date time.Time) (startTime string, endTime string) {
	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	endDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.Local)
	startTime = startDate.Format(FormatDate)
	endTime = endDate.Format(FormatDate)
	return
}

//输出开始时间和结束时间的日期
func GetBetweenDates(start_time, stop_time string) (args []string) {
	tm1, _ := time.Parse("2006-01-02", start_time)
	tm2, _ := time.Parse("2006-01-02", stop_time)

	sInt := tm1.Unix()
	eInt := tm2.Unix()
	if sInt == eInt {
		args = append(args, start_time)
		return
	}
	args = append(args, start_time)

	for {
		sInt += 86400
		st := time.Unix(sInt, 0).Format("2006-01-02")
		args = append(args, st)
		if sInt >= eInt {
			return
		}
	}
}

// 获取当天剩余的时间
func GetTodayRemainSec() uint64 {
	todayLast := time.Now().Format(FormatDay) + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation(FormatDate, todayLast, time.Local)
	remainSecond := uint64(todayLastTime.Unix() - time.Now().Local().Unix())
	return remainSecond
}

//获取昨天和今天
func YesterdayAndToday() []string {
	return []string{time.Now().AddDate(0, 0, -1).Format(FormatDay), time.Now().Format(FormatDay)}
}
