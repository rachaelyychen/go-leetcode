package main

import (
	"strconv"
	"strings"
)

func rotatedDigits(N int) int {
	var res int
	for i := 1; i <= N; i++ {
		if isGoodDigit(i) {
			res += 1
		}
	}
	return res
}

func isGoodDigit(x int) bool {
	s := strconv.Itoa(x)
	ok := strings.ContainsAny(s, "347")
	if ok {
		return false
	}
	ok = strings.ContainsAny(s, "2569")
	if ok {
		return true
	}
	return false
}


