package main

import "sort"

func repeatedNTimes(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return A[i]
		}
	}
	return 0
}


