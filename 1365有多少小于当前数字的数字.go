package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	m := make(map[int]int, 0)
	for i := range arr {
		if _, exists := m[arr[i]]; !exists {
			m[arr[i]] = i
		}
	}
	var res []int
	for i := range nums {
		index, _ := m[nums[i]]
		res = append(res, index)
	}
	return res
}


