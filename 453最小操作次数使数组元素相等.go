package main

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, min := 0, nums[0]
	for i := range nums {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	// ((n-1)*x + sum)/n = min+x
	return sum-min*len(nums)
}
