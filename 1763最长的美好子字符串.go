package main

import "strings"

func longestNiceSubstring(s string) string {
	var res string
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if checkNiceSubstring(i, j, arr) && len(res) < j-i+1 {
				res = string(arr[i : j+1])
			}
		}
	}
	return res
}

func checkNiceSubstring(from, to int, arr []byte) bool {
	// 1:缺小写字母 2:缺大写字母
	m := make(map[string]int, 0)
	for i := from; i <= to; i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if v, exists := m[strings.ToUpper(string(arr[i]))]; !exists {
				m[strings.ToUpper(string(arr[i]))] = 2
			} else if v == 1 {
				m[strings.ToUpper(string(arr[i]))] = 0
			}
		}
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			if v, exists := m[strings.ToUpper(string(arr[i]))]; !exists {
				m[strings.ToUpper(string(arr[i]))] = 1
			} else if v == 2 {
				m[strings.ToUpper(string(arr[i]))] = 0
			}
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
