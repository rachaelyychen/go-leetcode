package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
}

var arr []string

func init() {
	arr = make([]string, 31)
	arr[1] = "1"
	for i := 2; i <= 30; i++ {
		r := []rune(arr[i-1])
		res := ""
		for j := 0; j < len(r); j++ {
			count := 1
			x := r[j]
			for k := j + 1; k < len(r); k++ {
				if r[k] == x {
					count += 1
					j = k
				} else {
					j = k - 1
					break
				}
			}
			res += strconv.Itoa(count) + string(x)
		}
		arr[i] = res
	}
}

func countAndSay(n int) string {
	return arr[n]
}
