package main

func arraySign(nums []int) int {
	var t bool
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			t = !t
		}
	}
	if t {
		return -1
	}
	return 1
}
