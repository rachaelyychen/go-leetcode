package main

import (
	"strings"
)

func numDifferentIntegers(word string) int {
	arr := []byte(word)
	for i := range arr {
		if arr[i] < '0' || arr[i] > '9' {
			arr[i] = ' '
		}
	}
	strs := strings.Split(string(arr), " ")
	m := make(map[string]int, 0)
	for i := range strs {
		s := strs[i]
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		for strings.HasPrefix(s, "0") {
			s = strings.TrimPrefix(s, "0")
		}
		if cnt, exists := m[s]; !exists {
			m[s] = 1
		} else {
			m[s] = cnt + 1
		}
	}
	return len(m)
}
