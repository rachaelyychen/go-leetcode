package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	arr := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			t := nums[i] + nums[left] + nums[right]
			if t == 0 {
				arr = append(arr, []int{nums[i], nums[left], nums[right]})
				left += 1
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				right -= 1
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			} else if t < 0 {
				left += 1
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
			} else {
				right -= 1
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			}
		}
	}
	return arr
}
