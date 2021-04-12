package main

func specialArray(nums []int) int {
	res := -1
	for i := 1; i <= len(nums); i++ {
		var cnt int
		for j := range nums {
			if nums[j] >= i {
				cnt += 1
			}
		}
		if cnt == i {
			res = i
			break
		}
	}
	return res
}

