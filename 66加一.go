package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {
	x := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+x == 10 {
			digits[i] = 0
			x = 1
		} else {
			digits[i] += x
			x = 0
		}
	}
	if x > 0 {
		var newDigits []int
		newDigits = append(newDigits, x)
		newDigits = append(newDigits, digits...)
		return newDigits
	}
	return digits
}
