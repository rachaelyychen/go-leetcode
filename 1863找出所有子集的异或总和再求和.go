package main

func subsetXORSum(nums []int) int {
	var res int
	subsetXORSumDFS(nums, 0, 0, &res)
	return res
}

func subsetXORSumDFS(nums []int, index, sum int, res *int) {
	if index == len(nums) {
		*res += sum
		return
	}
	subsetXORSumDFS(nums, index+1, sum^nums[index], res)
	subsetXORSumDFS(nums, index+1, sum, res)
}
