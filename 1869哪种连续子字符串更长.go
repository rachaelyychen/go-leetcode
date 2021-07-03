package main

import "strings"

func checkZeroOnes(s string) bool {
	return getZeroOnes(s, '1') > getZeroOnes(s, '0')
}

func getZeroOnes(s string, b byte) int {
	var arr []string
	var res int
	if b == '0' {
		arr = strings.Split(s, "1")
	} else {
		arr = strings.Split(s, "0")
	}
	for i := range arr {
		if len(arr[i]) > res {
			res = len(arr[i])
		}
	}
	return res
}
