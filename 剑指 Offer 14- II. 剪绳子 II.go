package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/21/21 11:22 AM
**/

// 与上一题相比，这一题需要处理超出int64范围的大整数的比较大小，大整数来自长度积。

func init() {
	n := 1000
	mod := 1000000007
	dp = make([]dpBigInt, n+1)
	dp[1] = dpBigInt{[]int{1}, 1}
	dp[2] = dpBigInt{[]int{2}, 2}
	dp[3] = dpBigInt{[]int{3}, 3}
	for i := 4; i <= n; i++ {
		for k := 1; k <= i/2; k++ {
			var t dpBigInt
			t.arr = append(append(t.arr, dp[k].arr...), dp[i-k].arr...)
			t.b = (dp[k].b % mod * dp[i-k].b % mod) % mod
			if isLarger(t.arr, dp[i].arr) {
				dp[i] = t
			}
		}
	}
}

var dp []dpBigInt

type dpBigInt struct {
	arr []int
	b   int
}

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	return dp[n].b
}

func isLarger(x, y []int) bool {
	i, j := 0, 0
	res := float64(1)
	for {
		if i == len(x) && j == len(y) {
			return res > 1
		}
		if i == len(x) && res <= 1 {
			return false
		}
		if j == len(y) && res > 1 {
			return true
		}
		if i < len(x) {
			res *= float64(x[i])
			i += 1
		}
		if j < len(y) {
			res /= float64(y[j])
			j += 1
		}
	}
}
