package main

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	var t []int
	for i := range arr {
		t = append(t, arr[i])
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	m := make(map[int]int, 0)
	var d int
	for i := range t {
		if _, exists := m[t[i]]; !exists {
			m[t[i]] = i+1-d
		} else {
			d += 1
		}
	}
	var res []int
	for i := range arr {
		r, _ := m[arr[i]]
		res = append(res, r)
	}
	return res
}


