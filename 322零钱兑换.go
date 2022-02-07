package main

// dp[i][j]表示前i种硬币能凑成j的最少硬币数,dp[i][j]=min(dp[i-1][j],dp[i-1][j-c[i]*k]+k)
// 边界条件:dp[i][0]=0,dp[0][c[0]*k]=k

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < len(coins); i++ {
		dp = append(dp, make([]int, amount+1))
		for k := 1; k*coins[i] <= amount; k++ {
			if dp[i][k*coins[i]] == 0 {
				dp[i][k*coins[i]] = k
			} else if dp[i][k*coins[i]] > k {
				dp[i][k*coins[i]] = k
			}
		}
	}
	for i := 1; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if dp[i-1][j] > 0 {
				if dp[i][j] == 0 || dp[i][j] > dp[i-1][j] {
					dp[i][j] = dp[i-1][j]
				}
			}
			for k := 1; j-coins[i]*k >= 0; k++ {
				if dp[i-1][j-coins[i]*k] > 0 {
					if dp[i][j] == 0 || dp[i][j] > dp[i-1][j-coins[i]*k]+k {
						dp[i][j] = dp[i-1][j-coins[i]*k] + k
					}
				}
			}
		}
	}
	var res int
	for i := 0; i < len(coins); i++ {
		if dp[i][amount] == 0 {
			continue
		}
		if res == 0 {
			res = dp[i][amount]
		} else if dp[i][amount] < res {
			res = dp[i][amount]
		}
	}
	if res == 0 {
		return -1
	}
	return res
}
