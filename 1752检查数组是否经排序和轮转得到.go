package main

func check(nums []int) bool {
	var cnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			cnt += 1
		}
	}
	return cnt <= 1
}
