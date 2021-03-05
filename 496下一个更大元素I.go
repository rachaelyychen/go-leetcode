package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int
	for i := 0; i < len(nums1); i++ {
		var t int
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				t = j
				break
			}
		}
		var ok bool
		for j := t+1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				ok = true
				break
			}
		}
		if !ok {
			res = append(res, -1)
		}
	}
	return res
}
