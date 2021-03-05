package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return nums
	}
	tot := len(nums) * len(nums[0])
	if tot < r*c {
		return nums
	}
	var arr []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			arr = append(arr, nums[i][j])
		}
	}
	var res [][]int
	var index int
	for i := 0; i < r; i++ {
		var t []int
		for j := 0; j < c; j++ {
			t = append(t, arr[index])
			index += 1
		}
		res = append(res, t)
	}
	return res
}
