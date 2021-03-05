package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	fast, slow := squareSum(squareSum(n)), squareSum(n)
	for fast != slow {
		fast = squareSum(squareSum(fast))
		slow = squareSum(slow)
	}
	if fast == 1 {
		return true
	}
	return false
}

func squareSum(m int) int {
	var squareSum int
	for m != 0 {
		squareSum += (m%10)*(m%10)
		m /= 10
	}
	return squareSum
}
