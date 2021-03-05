package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
}

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return true
}
