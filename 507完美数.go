package main

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	factorSum, mid := 1, int(math.Sqrt(float64(num)))
	for i := 2; i <= mid; i++ {
		if num%i == 0 {
			factorSum = factorSum + i + num/i
		}
	}
	return factorSum==num
}
