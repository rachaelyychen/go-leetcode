package main

func sortedSquares(nums []int) []int {
	var arr []int
	i, j := 0, len(nums)-1
	for i < j {
		a, b := nums[i]*nums[i], nums[j]*nums[j]
		if a < b {
			arr = append(arr, b)
			j -= 1
		} else {
			arr = append(arr, a)
			i += 1
		}
	}
	arr = append(arr, nums[i]*nums[i])
	var res []int
	for i := len(arr)-1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}
