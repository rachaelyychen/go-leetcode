package main

func kLengthApart(nums []int, k int) bool {
	lastIndex := -1
	for i := range nums {
		if nums[i] == 1 {
			if lastIndex > -1 && i-lastIndex-1 < k {
				return false
			}
			lastIndex = i
		}
	}
	return true
}


