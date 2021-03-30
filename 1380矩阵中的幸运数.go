package main

import "sort"

func luckyNumbers (matrix [][]int) []int {
	rowMax, colMax := make([]int, len(matrix)), make([]int, len(matrix[0]))
	for i := range matrix {
		t := make([]int, len(matrix[i]))
		copy(t, matrix[i])
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		rowMax[i] = t[0]
	}
	for i := range matrix[0] {
		var t []int
		for j := range matrix {
			t = append(t, matrix[j][i])
		}
		sort.Slice(t, func(i, j int) bool {
			return t[i] > t[j]
		})
		colMax[i] = t[0]
	}
	var res []int
	for i := range matrix {
		for j := range matrix[i] {
			if rowMax[i] == colMax[j] {
				res = append(res, rowMax[i])
			}
		}
	}
	return res
}


