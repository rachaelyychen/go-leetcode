package main

import "strings"

func sortString(s string) string {
	var ans strings.Builder
	var arr [26]byte
	for _, v := range s {
		arr[v-'a'] ++
	}
	len_ := len(s)
	for len_ > 0 {
		for i := 0; i < 26; i++ {
			if arr[i] > 0 {
				ans.WriteByte('a' + byte(i))
				arr[i] -= 1
				len_ -= 1
			}
		}
		for i := 25; i >= 0; i-- {
			if arr[i] > 0 {
				ans.WriteByte('a' + byte(i))
				arr[i] -= 1
				len_ -= 1
			}
		}
	}
	return ans.String()
}
