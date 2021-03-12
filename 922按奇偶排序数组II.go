package main

func sortArrayByParityII(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, 0
	for k := 0; k < len(res); k++ {
		if k % 2 == 0 {
			for i < len(nums) && nums[i] % 2 != 0 {
				i += 1
			}
			if i < len(nums) {
				res[k] = nums[i]
				i += 1
			}
		} else {
			for j < len(nums) && nums[j] % 2 == 0 {
				j += 1
			}
			if j < len(nums) {
				res[k] = nums[j]
				j += 1
			}
		}
	}
	return res
}


