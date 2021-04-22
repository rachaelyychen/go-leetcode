package main

import "strings"

func maximumTime(time string) string {
	arr := []byte(strings.ReplaceAll(time, ":", ""))
	var res string
	for i := 0; i < 4; i++ {
		if arr[i] == '?' {
			switch (i) {
			case 0:
				if arr[1] >= '4' && arr[1] <= '9' {
					arr[0] = '1'
				} else {
					arr[0] = '2'
				}
			case 1:
				if arr[0] == '2' {
					arr[1] = '3'
				} else {
					arr[1] = '9'
				}
			case 2:
				arr[2] = '5'
			case 3:
				arr[3] = '9'
			}
		}
		res += string(arr[i])
		if i == 1 {
			res += ":"
		}
	}
	return res
}
