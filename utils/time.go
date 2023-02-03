package scyna_utils

import (
	"time"
)

const microSecondPerDay = 24 * 60 * 60 * 1000000

func GetDayByTime(time time.Time) int {
	return int(time.UnixMicro() / microSecondPerDay)
}

func GetMinuteByTime(time time.Time) int64 {
	return time.Unix() / 60
}

func GetHourByTime(time time.Time) int64 {
	return time.Unix() / (60 * 60)
}
