package main

import "strings"

func makeGood(s string) string {
	arr, res := []byte(s), make([]byte, 0)
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && !isGoodChars(arr[i], arr[i+1]) {
			i += 1
			continue
		}
		res = append(res, arr[i])
	}
	if string(res) == s {
		return s
	}
	return makeGood(string(res))
}
func isGoodChars(x, y byte) bool {
	if x != y && (strings.ToLower(string(x)) == string(y) || strings.ToLower(string(y)) == string(x)) {
		return false
	}
	return true
}
