package main

import "strings"

func truncateSentence(s string, k int) string {
	arr := []byte(s)
	var res []byte
	var cnt int
	for i := 0; i <= len(arr); i++ {
		if i == len(arr) || cnt == k {
			break
		}
		res = append(res, arr[i])
		if arr[i] == ' ' {
			cnt += 1
		}
	}
	return strings.TrimSuffix(string(res), " ")
}
