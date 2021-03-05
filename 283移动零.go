package main

func moveZeroes(nums []int)  {
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i] = nums[j]
			nums[j] = 0
			i += 1
			j += 1
		} else if nums[i] == 0 && nums[j] == 0 {
			j += 1
		} else {
			i += 1
			j += 1
		}
	}
}
