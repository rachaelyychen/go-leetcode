package main

func findShortestSubArray(nums []int) int {
	arr := make([]int, 50000)
	for i := range nums {
		arr[nums[i]] += 1
	}
	var maxCnt int
	for i := range arr {
		if arr[i] > maxCnt {
			maxCnt = arr[i]
		}
	}
	res := len(nums)
	for x := range arr {
		if arr[x] == maxCnt {
			first, last := -1, -1
			for i := range nums {
				if nums[i] == x && first == -1 {
					first = i
				}
				if nums[i] == x && first > -1 {
					last = i
				}
			}
			if last - first + 1 < res {
				res = last -first + 1
			}
		}
	}
	return res
}
