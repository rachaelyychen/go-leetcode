package main

func increasingTriplet(nums []int) bool {
	leftMinArr := make([]int, len(nums))
	leftMin, rightMax := nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		leftMinArr[i] = leftMin
		if nums[i] < leftMin {
			leftMin = nums[i]
		}
	}
	for i := len(nums) - 2; i > 0; i-- {
		if leftMinArr[i] < nums[i] && nums[i] < rightMax {
			return true
		}
		if nums[i] > rightMax {
			rightMax = nums[i]
		}
	}
	return false
}
