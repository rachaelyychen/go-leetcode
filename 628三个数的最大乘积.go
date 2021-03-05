package main

import "sort"

func maximumProduct(nums []int) int {
	arr1, arr2 := make([]int, 0), make([]int, 0)
	for i := range nums {
		if nums[i] < 0 {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	// 负数正序
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	// 正数倒序
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] > arr2[j]
	})
	var res int
	if len(nums) == 3 {
		res = nums[0] * nums[1] * nums[2]
	} else {
		if len(arr1) == 0 {
			res = arr2[0] * arr2[1] * arr2[2]
		} else {
			if len(arr1) > 1 {
				if len(arr2) > 0 {
					res = arr1[0] * arr1[1] * arr2[0]
					if len(arr2) > 2 && res <  arr2[0] * arr2[1] * arr2[2] {
						res =  arr2[0] * arr2[1] * arr2[2]
					}
				} else {
					res = arr1[len(arr1)-1]*arr1[len(arr1)-2]*arr1[len(arr1)-3]
				}
			} else {
				res = arr2[0] * arr2[1] * arr2[2]
			}
		}
	}
	return res
}
