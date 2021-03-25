package main

import "strconv"

func maximum69Number (num int) int {
	arr := []byte(strconv.Itoa(num))
	for i := range arr {
		if arr[i] == '6' {
			arr[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(arr))
	return res
}


