package main

import "fmt"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/20/21 4:16 PM
**/

// 动态规划的三大特点：1、求最优解 2、大问题可以分解成若干小问题，小问题之间有重叠 3、大问题的最优解依赖小问题的最优解。
// 第一种解法，dp[i][j]表示长为i的绳子剪成j段能得到的最大长度积，边界条件：dp[i][i]=1
// 动态规划方程：dp[i][j]=max(dp[i-1][j-1],dp[i-2][j-1]*2...dp[i-k][j-1]*k) (k=1,2,...)，满足条件i-k>=j-1,1<i<=n,1<j<=n
// 第二种解法，相比第一种解法，降低时间复杂度。
// dp[i]表示长度为i的绳子能得到的最大长度积，边界条件：dp[1]=1, dp[2]=2, dp[3]=3
// 动态规划方程：dp[i]=max(dp[k]*dp[i-k])，满足条件4=<i<=n,0<k<i
// 需要注意的是，对于长度小于4的绳子，题目要求是强制剪至少一次，但在长绳子剪成短绳子的分解中可以不剪，所以dp数组的元素有所不同。

func main() {
	fmt.Println(cuttingRope(8), cuttingRope1(8))
}

func cuttingRope(n int) int {
	dp := [60][60]int{}
	for i := 2; i <= n; i++ {
		dp[i][i] = 1
		if i%2 == 0 {
			dp[i][2] = (i / 2) * (i / 2)
		} else {
			dp[i][2] = (i / 2) * (i/2 + 1)
		}
	}
	for i := 3; i <= n; i++ {
		for j := 3; j <= n; j++ {
			if i <= j {
				continue
			}
			var t int
			for k := 1; i-k >= j-1; k++ {
				if t < dp[i-k][j-1]*k {
					t = dp[i-k][j-1] * k
				}
			}
			dp[i][j] = t
		}
	}
	var res int
	for i := 2; i <= n; i++ {
		if res < dp[n][i] {
			res = dp[n][i]
		}
	}
	return res
}

func cuttingRope1(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		for k := 1; k <= i/2; k++ {
			t := dp[k] * dp[i-k]
			if dp[i] < t {
				dp[i] = t
			}
		}
	}
	return dp[n]
}
