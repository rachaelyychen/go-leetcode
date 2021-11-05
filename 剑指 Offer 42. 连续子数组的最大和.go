package main

// 时间复杂度O(n)的解法：遍历数组，res记录最大和，sum记录之前的累加和，
// 如果累加和小于0，由于只会让当前数字变小，于是抛弃之前的累加和，而从当前数字开始累加，这里需要更新最大和；
// 如果累加和不小于0，但是当前数字小于0，会让累加和变小，所以这个累加和可能就是最大和，需要先更新最大和，再做累加。

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := res
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
			if sum > res {
				res = sum
			}
			continue
		}
		if nums[i] < 0 {
			if sum > res {
				res = sum
			}
		}
		sum += nums[i]
	}
	if sum > res {
		res = sum
	}
	return res
}
