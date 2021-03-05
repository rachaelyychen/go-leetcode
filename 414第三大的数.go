package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
}

func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	order := 1
	for i := 1; i < len(nums); i++ {
		for i < len(nums) && nums[i] == nums[i-1] {
			i += 1
		}
		if i < len(nums) {
			order += 1
			if order == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}
