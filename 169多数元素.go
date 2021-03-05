package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}

func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var x int
	if len(nums) % 2 == 0 {
		x = len(nums) / 2 - 1
	} else {
		x = len(nums) / 2
	}
	return nums[x]
}
