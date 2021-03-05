package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSegments("hello, world,     a"))
}

func countSegments(s string) int {
	arr := []rune(s)
	var res int
	var inWord bool
	for i := 0; i < len(arr); i++ {
		if inWord == false && arr[i] != ' ' {
			inWord = true
		}
		if inWord == true && (i == len(arr)-1 || arr[i] == ' ') {
			inWord = false
			res += 1
		}
	}
	return res
}
