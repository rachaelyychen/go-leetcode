package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
	if len(s) > 0 {
		arr := []byte(s)
		for i := 1; i < len(s); i++ {
			subStr := string(arr[:i])
			if len(s)%len(subStr) == 0 && strings.Repeat(subStr, len(s)/len(subStr)) == s {
				// fmt.Println(subStr)
				return true
			}
		}
	}
	return false
}
