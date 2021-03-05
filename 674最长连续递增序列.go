package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, from := 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if i-from+1 > res {
				res = i-from+1
			}
		} else {
			from = i
		}
	}
	return res
}
