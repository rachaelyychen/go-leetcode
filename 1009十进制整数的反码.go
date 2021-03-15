package main

import "math"

func bitwiseComplement(N int) int {
	var arr []int
	if N == 0 {
		arr = append(arr, 0)
	} else {
		for N > 0 {
			arr = append(arr, N%2)
			N /= 2
		}
	}
	var res int
	for i := range arr {
		if arr[i] == 0 {
			res += int(math.Pow(2, float64(i)))
		}
	}
	return res
}


