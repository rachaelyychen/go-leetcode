package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumber([]int{4,1,2,1,2}))
}

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := range nums {
		if i % 2 > 0 || i+1<len(nums) && nums[i] == nums[i+1] {
			continue
		}
		return nums[i]
	}
	return 0
}
