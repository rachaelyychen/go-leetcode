package main

import "sort"

func maxProductDifference(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}
