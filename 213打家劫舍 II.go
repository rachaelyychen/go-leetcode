package main

// dp[i]表示从第0个房间到第i个房间能获得的最大值,dp[i]=max(dp[i-2]+nums[i],dp[i-1])
// 边界条件:dp[0]=nums[0],dp[1]=dp[0] 需要注意最后一个房间不能选
// 边界条件:dp[0]=0,dp[1]=nums[1]

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var res int
	if len(nums) == 2 {
		res = nums[0]
		if res < nums[1] {
			res = nums[1]
		}
		return res
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = dp[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if i < len(nums)-1 {
			if dp[i] < dp[i-2]+nums[i] {
				dp[i] = dp[i-2] + nums[i]
			}
		}
	}
	res = dp[len(nums)-1]
	dp[0] = 0
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if dp[i] < dp[i-2]+nums[i] {
			dp[i] = dp[i-2] + nums[i]
		}
	}
	if res < dp[len(nums)-1] {
		res = dp[len(nums)-1]
	}
	return res
}
