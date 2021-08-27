package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	arr, res := make([][]int, 0), make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 || i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right && left > i && right < len(nums) {
			t := nums[i] + nums[left] + nums[right]
			if t < 0 {
				left += 1
			} else if t > 0 {
				right -= 1
			} else {
				arr = append(arr, []int{nums[i], nums[left], nums[right]})
				left += 1
				right -= 1
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		if arr[i][1] != arr[j][1] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][2] < arr[j][2]
	})
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i][0] != arr[i-1][0] || arr[i][1] != arr[i-1][1] || arr[i][2] != arr[i-1][2] {
			res = append(res, arr[i])
		}
	}
	return res
}
