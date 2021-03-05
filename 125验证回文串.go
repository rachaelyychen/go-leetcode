package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome125("a"))
}

func isPalindrome125(s string) bool {
	if len(s) == 0 {
		return true
	}
	arr := []rune(strings.ToLower(s))
	var ns string
	for i := 0; i < len(arr); i++ {
		if unicode.IsDigit(arr[i]) || unicode.IsLetter(arr[i]) {
			ns += string(arr[i])
		}
	}
	var rns string
	for i := len(arr)-1; i >= 0; i-- {
		if unicode.IsDigit(arr[i]) || unicode.IsLetter(arr[i]) {
			rns += string(arr[i])
		}
	}
	return strings.EqualFold(ns, rns)
}
