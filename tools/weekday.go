package tools

import "time"

func WeekDayNorm() int {
	curWeekDay := time.Now().Weekday()
	if curWeekDay == 0 {
		return 7
	}
	return int(curWeekDay)
}
