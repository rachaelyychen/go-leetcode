package main

func getMinDistance(nums []int, target int, start int) int {
	res := -1
	for i := range nums {
		if nums[i] == target {
			var t int
			if i < start {
				t = start - i
			} else {
				t = i - start
			}
			if res < 0 || t < res {
				res = t
			}
		}
	}
	return res
}
