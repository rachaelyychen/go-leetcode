package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{10,2,5,10,9,1,1,4,3,7}))
}

func findDisappearedNumbers(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res []int
	// 1 1 2 3 4 5 7 9 10 10
	// 1 2 3 4 5 6 7 8 9 10
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			var ok bool
			if nums[i] < i+1 {
				for j := i+1; j < len(nums); j++ {
					if nums[j] == i+1 {
						nums[j] = 0
						ok = true
						break
					}
				}
			} else {
				for j := 0; j < i; j++ {
					if nums[j] == i+1 {
						nums[j] = 0
						ok = true
						break
					}
				}
			}
			if !ok {
				res = append(res, i+1)
			}
		}
	}
	return res
}
