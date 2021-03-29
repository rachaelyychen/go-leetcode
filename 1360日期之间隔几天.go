package main

import (
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)
	return int(math.Abs(t1.Sub(t2).Hours())) / 24
}
