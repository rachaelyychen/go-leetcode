package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}

func summaryRanges(nums []int) []string {
	var res []string
	for i := 0; i < len(nums); i++ {
		s := strconv.Itoa(nums[i])
		var t string
		for j := i+1; j < len(nums); j++ {
			if nums[j] == nums[j-1]+1 {
				t = "->"+strconv.Itoa(nums[j])
				i = j
			} else {
				i = j-1
				break
			}
		}
		s += t
		res = append(res, s)
	}
	return res
}
