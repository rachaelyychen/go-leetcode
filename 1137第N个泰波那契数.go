package main

var tribonacciArr []int

func tribonacci(n int) int {
	return tribonacciArr[n]
}

func init() {
	for i := 0; i <= 37; i++ {
		if i == 0 {
			tribonacciArr[i] = 0
		}
		if i == 1 || i == 2 {
			tribonacciArr[i] = 1
		}
		if i > 2 {
			tribonacciArr[i] = tribonacciArr[i-1] + tribonacciArr[i-2] + tribonacciArr[i-3]
		}
	}
}


