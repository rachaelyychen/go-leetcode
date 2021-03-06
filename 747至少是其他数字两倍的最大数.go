package main

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxIndex, max := -1, -1
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}
	for i := range nums {
		if i != maxIndex && nums[i] != max && 2*nums[i] > max {
			return -1
		}
	}
	return maxIndex
}
