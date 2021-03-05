package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersection2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection2(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	var res []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i += 1
			j += 1
		} else if nums1[i] < nums2[j] {
			i += 1
		} else {
			j += 1
		}
	}
	return res
}
