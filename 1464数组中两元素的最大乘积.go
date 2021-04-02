package main

import "sort"

func maxProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
