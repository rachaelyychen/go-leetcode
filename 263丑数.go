package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num % 5 == 0 {
		num /= 5
	}
	for num % 3 == 0 {
		num /= 3
	}
	for num % 2 == 0 {
		num /= 2
	}
	return num == 1
}
