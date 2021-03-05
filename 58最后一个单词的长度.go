package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("b   a    "))
}

func lengthOfLastWord(s string) int {
	for strings.HasSuffix(s, " ") {
		s = strings.TrimSuffix(s, " ")
	}
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}
