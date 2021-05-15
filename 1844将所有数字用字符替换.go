package main

import "strconv"

func replaceDigits(s string) string {
	arr := []byte(s)
	for i := range arr {
		if i%2 > 0 {
			t, _ := strconv.Atoi(string(arr[i]))
			arr[i] = byte(int(arr[i-1]) + t)
		}
	}
	return string(arr)
}
