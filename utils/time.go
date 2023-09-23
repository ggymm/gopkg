package utils

import "time"

func Now() int64 {
	return time.Now().UnixMilli()
}

func MinuteToSecond(minute int) int {
	return minute * 60
}

func HourToSecond(hour int) int {
	return MinuteToSecond(hour * 60)
}

func DayToSecond(day int) int {
	return HourToSecond(day * 24)
}

func MonthToSecond(month int) int {
	return DayToSecond(month * 30)
}

func YearToSecond(year int) int {
	return MonthToSecond(year * 12)
}

func Format(format string, t time.Time) string {
	return t.Format(format)
}

func FormatMilli(format string, t int64) string {
	return Format(format, time.UnixMilli(t))
}
