package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reformatNumber("123 4-567"))
}

func reformatNumber(number string) string {
	number = strings.ReplaceAll(strings.ReplaceAll(number, " ", ""), "-", "")
	var res string
	arr := []byte(number)
	cnt := len(arr) / 3
	if len(arr)%3 > 0 {
		cnt += 1
	}
	if len(arr)%3 == 1 {
		cnt -= 2
	}
	i := 0
	for i = 0; i <= len(arr); i++ {
		if cnt == 0 || i == len(arr) {
			break
		}
		res += string(arr[i])
		if (i+1)%3 == 0 && i < len(arr)-1 {
			res += "-"
			cnt -= 1
		}
	}
	if i != len(arr) {
		res += string(arr[len(arr)-4]) + string(arr[len(arr)-3]) + "-" + string(arr[len(arr)-2]) + string(arr[len(arr)-1])
	}
	return res
}
