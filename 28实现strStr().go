package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
