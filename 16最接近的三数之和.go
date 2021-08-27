package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var t, res int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			x := nums[i] + nums[left] + nums[right] - target
			if x == 0 {
				return target
			} else if x < 0 {
				if t == 0 || t > -x {
					t = -x
					res = nums[i] + nums[left] + nums[right]
				}
				left += 1
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
			} else {
				if t == 0 || t > x {
					t = x
					res = nums[i] + nums[left] + nums[right]
				}
				right -= 1
				for right > left && nums[right] == nums[right+1] {
					right -= 1
				}
			}
		}
	}
	return res
}
