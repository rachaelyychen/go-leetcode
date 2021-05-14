package main

func maxAscendingSum(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		t, s := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > t {
				t = nums[j]
				s += nums[j]
			} else {
				break
			}
		}
		if s > res {
			res = s
		}
	}
	return res
}
