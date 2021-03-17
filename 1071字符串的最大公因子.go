package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	var s, t string
	if len(str1) < len(str2) {
		s = str1
		t = str2
	} else {
		s = str2
		t = str1
	}
	for i := len(s); i > 0; i-- {
		if len(s)%i == 0 && len(t)%i == 0 {
			x := s[:i]
			if strings.Repeat(x, len(s)/i) == s && strings.Repeat(x, len(t)/i) == t {
				return x
			}
		}
	}
	return ""
}


