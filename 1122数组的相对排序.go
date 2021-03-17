package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var t1 []int
	var res []int
	t2 := make([]int, 1001)
	for i := range arr2 {
		t2[arr2[i]] = i+1
	}
	for i := range arr1 {
		if t2[arr1[i]] == 0 {
			t1 = append(t1, arr1[i])
		} else {
			res = append(res, arr1[i])
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return t2[res[i]] < t2[res[j]]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	res = append(res, t1...)
	return res
}


