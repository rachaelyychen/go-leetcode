package main

import "math"

// dp[i]表示和为数字i的完全平方数的最少数量,dp[i]=min(dp[i-j*j])+1
// 边界条件:dp[0]=0

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if t > dp[i-j*j] {
				t = dp[i-j*j]
			}
		}
		dp[i] = t + 1
	}
	return dp[n]
}
