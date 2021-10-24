package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 12:50 AM
**/

// 考虑x是非0和0、n是非负数和负数这些情况，以及使用二分思想将幂运算的时间复杂度降低到O(logn)。

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var m int
	if n < 0 {
		m = -n
	} else {
		m = n
	}
	res := myUIntPow(x, uint(m))
	if n < 0 {
		return 1 / res
	}
	return res
}

func myUIntPow(x float64, m uint) float64 {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return x
	}
	t := myUIntPow(x, m/2)
	if m%2 == 0 {
		return t * t
	} else {
		return t * t * x
	}
}
