package main

import (
	"strconv"
)

func getLucky(s string, k int) int {
	arr := []byte(s)
	var t string
	for i := range arr {
		t += strconv.Itoa(int(arr[i]-'a') + 1)
	}
	arr = []byte(t)
	var res int
	for i := range arr {
		a, _ := strconv.Atoi(string(arr[i]))
		res += a
	}
	for k > 1 {
		k -= 1
		var n int
		for res > 0 {
			n += res % 10
			res /= 10
		}
		res = n
	}
	return res
}
