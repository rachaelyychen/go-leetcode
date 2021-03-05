package main

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var res string
	for i := range arr {
		res += reverseWordString(arr[i])
		if i < len(arr)-1 {
			res += " "
		}
	}
	return res
}

func reverseWordString(s string) string {
	arr := []byte(s)
	var res string
	for i := len(arr)-1; i >= 0; i-- {
		res += string(arr[i])
	}
	return res
}















