package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(2))
}

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	var tot int
	for i := 1; i <= n; i++ {
		if n-tot < i {
			return i-1
		}
		tot += i
	}
	return 0
}
