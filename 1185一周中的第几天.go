package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	t, _ := time.ParseInLocation("2006-01-02",
		fmt.Sprintf("%04d-%02d-%02d", year, month, day), time.Local)
	return t.Weekday().String()
}
