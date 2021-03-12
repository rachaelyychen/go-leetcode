package main

import "sort"

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i+2 < len(nums) && nums[i+1]+nums[i+2] > nums[i] {
			return nums[i+1]+nums[i+2]+nums[i]
		}
	}
	return 0
}


