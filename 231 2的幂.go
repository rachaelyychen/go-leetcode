package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(5))
}

func isPowerOfTwo(n int) bool {
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	if n == 1 {
		return true
	}
	return false
}
