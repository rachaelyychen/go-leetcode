package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}

func sortSentence(s string) string {
	arr1, arr2 := strings.Split(s, " "), [10]string{}
	for i := range arr1 {
		b := []byte(arr1[i])
		arr2[b[len(b)-1] - '0' - 1] = string(b[:len(b)-1])
	}
	var res string
	for i := range arr2 {
		if len(arr2[i]) > 0 {
			if i > 0 {
				res += " "
			}
			res += arr2[i]
		}
	}
	return res
}
