package main

import (
	"sort"
)

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] != ranges[j][0] {
			return ranges[i][0] < ranges[j][0]
		}
		return ranges[i][1] < ranges[j][1]
	})
	var t [][]int
	index := -1
	for i := range ranges {
		if index == -1 || ranges[i][0] > t[index][1] {
			t = append(t, ranges[i])
			index += 1
			continue
		}
		if ranges[i][1] > t[index][1] {
			t[index][1] = ranges[i][1]
		}
	}
	res := make([]bool, right-left+1)
	index = 0
	for i := left; i <= right; i++ {
		for j := index; j < len(t); j++ {
			if t[j][0] <= i && i <= t[j][1] {
				res[i-left] = true
				index = j
				break
			}
		}
	}
	for i := range res {
		if res[i] == false {
			return false
		}
	}
	return true
}
