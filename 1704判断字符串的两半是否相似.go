package main

import "strings"

func halvesAreAlike(s string) bool {
	arr := []byte(s)
	var a, b string
	for i := 0; i < len(arr)/2; i++ {
		a += string(arr[i])
	}
	b = strings.TrimPrefix(s, a)
	var cnt int
	for _, x := range []byte(a) {
		if isVowel(x) {
			cnt += 1
		}
	}
	for _, x := range []byte(b) {
		if isVowel(x) {
			cnt -= 1
		}
	}
	return cnt == 0
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' || x == 'A' || x == 'E' || x == 'I' || x == 'O' || x == 'U'
}
