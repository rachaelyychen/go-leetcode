package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("a good a good boy a good student", "a", "good"))
}

func findOcurrences(text string, first string, second string) []string {
	var res []string
	s := strings.Split(text, " ")
	for i := 0; i < len(s)-2; i++ {
		if s[i] == first && s[i+1] == second {
			res = append(res, s[i+2])
		}
	}
	return res
}
