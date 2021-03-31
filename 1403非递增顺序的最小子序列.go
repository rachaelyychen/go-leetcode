package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum1, sum2 := 0, 0
	for i := range nums {
		sum1 += nums[i]
	}
	for i := range nums {
		sum2 += nums[i]
		if sum2 > sum1-sum2 {
			return nums[:i+1]
		}
	}
	return nil
}


