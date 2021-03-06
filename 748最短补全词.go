package main

import (
	"bytes"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	plateArr := []byte(licensePlate)
	m := make(map[byte]int, 0)
	for i := range plateArr {
		if plateArr[i] >= 'a' && plateArr[i] <= 'z' || plateArr[i] >= 'A' && plateArr[i] <= 'Z' {
			b := bytes.ToLower([]byte{plateArr[i]})
			if v, exists := m[b[0]]; !exists {
				m[b[0]] = 1
			} else {
				m[b[0]] = v + 1
			}
		}
	}
	var res string
	for i := range words {
		ok := true
		for b, cnt := range m {
			if strings.Count(words[i], string(b)) < cnt {
				ok = false
				break
			}
		}
		if ok {
			if res == "" || len(res) > len(words[i]) {
				res = words[i]
			}
		}
	}
	return res
}
