package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		res += int(num%2)
		num >>= 1
	}
	return res
}
