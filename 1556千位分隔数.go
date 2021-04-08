package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(thousandSeparator(0))
}

func thousandSeparator(n int) string {
	res, arr := "", make([]string, 0)
	if n == 0 {
		arr = append(arr, "0")
	} else {
		for n > 0 {
			arr = append(arr, strconv.Itoa(n%10))
			n /= 10
		}
	}
	var t int
	for i := 0; i < len(arr); i++ {
		res = arr[i] + res
		t += 1
		if t == 3 {
			res = "." + res
			t = 0
		}
	}
	return strings.TrimPrefix(res, ".")
}
