package main

func decompressRLElist(nums []int) []int {
	var res []int
	for i := range nums {
		if i%2 == 0 && i+1 < len(nums) {
			for j := 0; j < nums[i]; j++ {
				res = append(res, nums[i+1])
			}
		}
	}
	return res
}
