package timeStamp

import (
	"time"
)

const (
	sun = iota
	mon
	tue
	wed
	thu
	fri
	sat
	unixDay = 86400
)

func daysToAdd(day int, current int) int {
	return ((day - current + 7) % 7)
}

func GetUnixTimeStamp() (int, int, error) {
	var start int
	var end int

	unixNow := time.Now().Unix()
	unixInt := int(unixNow)

	today := time.Now().Weekday()
	todayInt := int(today)

	dta := daysToAdd(mon, todayInt)

	secsElapsed := unixInt % unixDay

	start = (unixInt + (dta * unixDay) - secsElapsed)
	end = (unixInt + ((dta + 7) * unixDay) - secsElapsed)

	return start, end, nil
}
