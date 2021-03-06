package main

func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] < target && (i == len(nums)-1 || i+1 < len(nums) && nums[i+1] > target) {
			return -1
		}
	}
	return -1
}
