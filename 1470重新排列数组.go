package main

func shuffle(nums []int, n int) []int {
	// 0 ~ n-1, n ~ 2n-1
	var res []int
	i, j := 0, n
	for i < n && j < 2*n {
		res = append(res, nums[i], nums[j])
		i += 1
		j += 1
	}
	return res
}


