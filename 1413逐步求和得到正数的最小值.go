package main

func minStartValue(nums []int) int {
	sum, minSum := 0, 100*100
	for i := range nums {
		sum += nums[i]
		if sum < minSum {
			minSum = sum
		}
	}
	var res int
	if minSum > 0 {
		res = 1
	} else {
		res = 1 - minSum
	}
	return res
}
