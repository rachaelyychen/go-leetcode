package main

import "fmt"

func main() {
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	t, c := num, 0
	for t > 0 {
		t >>= 1
		c = (c << 1) + 1
	}
	return num ^ c
}
