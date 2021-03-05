package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(5))
}

func trailingZeroes(n int) int {
	var res int
	for n >= 5 {
		res += n/5
		n /= 5
	}
	return res
}
