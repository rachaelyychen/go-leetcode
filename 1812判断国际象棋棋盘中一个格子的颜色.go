package main

import "strconv"

func squareIsWhite(coordinates string) bool {
	arr := []byte(coordinates)
	t1 := arr[0] - 'a'
	t2, _ := strconv.Atoi(string(arr[1]))
	if t1%2 == 0 && t2%2 == 0 || t1%2 == 1 && t2%2 == 1 {
		return true
	}
	return false
}
