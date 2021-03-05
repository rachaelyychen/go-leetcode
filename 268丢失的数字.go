package main

import "sort"

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] - nums[i] == 2 {
			return nums[i]+1
		}
	}
	// 0 或者 n
	if nums[len(nums)-1] == len(nums) {
		return 0
	} else {
		return len(nums)
	}
}
