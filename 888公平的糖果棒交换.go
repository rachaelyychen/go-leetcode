package main

func fairCandySwap(A []int, B []int) []int {
	// 2*(x - y) = a - b
	a, b := 0, 0
	for i := range A {
		a += A[i]
	}
	for i := range B {
		b += B[i]
	}
	for i := range A {
		for j := range B {
			if 2*(A[i]-B[j]) == a-b {
				return []int{A[i], B[j]}
			}
		}
	}
	return nil
}
