package main

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	// 优先给所有负数取反，从小到大，如果还有剩余，偶数没关系，奇数给当前最小值
	// 如果没有负数，偶数没关系，奇数给当前最小值
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	var minNum int
	for i := range A {
		if A[i] < 0 && K > 0 {
			A[i] = -A[i]
			K -= 1
		}
		if i == 0 {
			minNum = A[i]
		} else if minNum > A[i] {
			minNum = A[i]
		}
	}
	var res int
	for i := range A {
		res += A[i]
	}
	if K > 0 && K%2 == 1 {
		res -= minNum
		res += -minNum
	}
	return res
}
