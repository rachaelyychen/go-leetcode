package main

import "strings"

var months = map[string]string {
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func formatDay(day string) string {
	result := day[:len(day) - 2]
	if len(result) == 1 {
		return "0" + result
	}
	return result
}

func reformatDate(date string) string {
	parts := strings.Split(date, " ")
	day, month, year := parts[0], parts[1], parts[2]
	return strings.Join([]string{year, months[month], formatDay(day)}, "-")
}

