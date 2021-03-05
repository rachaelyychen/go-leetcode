package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	z, res := x^y, 0
	for z != 0 {
		res += z & 1
		z >>= 1
	}
	return res
}
