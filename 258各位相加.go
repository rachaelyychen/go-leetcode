package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	for num > 9 {
		x := num
		num = 0
		for x > 0 {
			num += x%10
			 x /= 10
		}
	}
	return num
}
