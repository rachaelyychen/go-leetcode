package main

// dp[i][j]表示前i个数字是否存在和为j的方案,dp[i][j]=dp[i-1][j] || dp[i-1][j-nums[i]]
// 边界条件:dp[i][0]=true,dp[i][nums[i]]=true

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, 0)
	for i := 0; i < len(nums); i++ {
		dp = append(dp, make([]bool, target+1))
		dp[i][0] = true
		if nums[i] <= target {
			dp[i][nums[i]] = true
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if dp[i][j] == false && j-nums[i] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i]]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if dp[i][target] == true {
			return true
		}
	}
	return false
}
