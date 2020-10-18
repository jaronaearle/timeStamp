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

	timeNow := time.Now().Unix()
	unixNow := int(timeNow)

	tt := time.Now().Weekday()
	it := int(tt)

	dta := daysToAdd(mon, it)

	secs := unixNow % unixDay

	start = (unixNow + (dta * unixDay) - secs)
	end = (unixNow + ((dta + 7) * unixDay) - secs)

	return start, end, nil
}
