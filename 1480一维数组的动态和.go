package main

func runningSum(nums []int) []int {
	var res []int
	var sum int
	for i := range nums {
		sum += nums[i]
		res = append(res, sum)
	}
	return res
}


