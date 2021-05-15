package main

func minOperations(nums []int) int {
	var res int
	if len(nums) <= 1 {
		return res
	}
	for i := range nums {
		if i+1 < len(nums) {
			if nums[i+1] < nums[i] {
				res += nums[i] - nums[i+1] + 1
				nums[i+1] = nums[i] + 1
			} else if nums[i+1] == nums[i] {
				res += 1
				nums[i+1] = nums[i] + 1
			}
		}
	}
	return res
}
