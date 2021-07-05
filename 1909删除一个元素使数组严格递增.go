package main

func canBeIncreasing(nums []int) bool {
	var cnt int
	for i := range nums {
		if i+1 < len(nums) && nums[i] >= nums[i+1] {
			cnt += 1
			if cnt > 1 {
				return false
			}
			if i-1 >= 0 && nums[i-1] >= nums[i+1] && i+2 < len(nums) && nums[i] >= nums[i+2] {
				return false
			}
		}
	}
	return true
}
