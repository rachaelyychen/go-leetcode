package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(701))
}

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}
	var res string
	for n > 0 {
		n -= 1
		res = string(n%26+'A') + res
		n /= 26
	}
	return res
}
