package main

import "sort"

func heightChecker(heights []int) int {
	var t []int
	for i := range heights {
		t = append(t, heights[i])
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	var res int
	for i := range t {
		if t[i] != heights[i] {
			res += 1
		}
	}
	return res
}


