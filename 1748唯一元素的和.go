package main

func sumOfUnique(nums []int) int {
	var res int
	arr := make([]int, 105)
	for i := range nums {
		arr[nums[i]] += 1
	}
	for i := range arr {
		if arr[i] == 1 {
			res += i
		}
	}
	return res
}


