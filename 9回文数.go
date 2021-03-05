package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isPalindrome(88))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s1 := ""
	t := x
	for ; t != 0; {
		s1 += strconv.Itoa(t%10)
		t /= 10
	}
	s2 := strconv.Itoa(x)
	if strings.EqualFold(s1, s2) {
		return true
	}
	return false
}
