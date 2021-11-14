package main

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		t := nums[left] + nums[right]
		if t == target {
			return []int{nums[left], nums[right]}
		} else if t > target {
			right -= 1
		} else {
			left += 1
		}
	}
	return []int{}
}
