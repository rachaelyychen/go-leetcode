package main

import "sort"

func smallestRangeI(A []int, K int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	res := A[len(A)-1] - A[0] - 2*K
	if res < 0 {
		res = 0
	}
	return res
}


