package main

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	minNum, maxNum := 2, arr[2]-arr[0]-2
	if arr[2]-arr[1] == 1 && arr[1]-arr[0] == 1 {
		minNum = 0
	} else if arr[2]-arr[1] <= 2 || arr[1]-arr[0] <= 2 {
		minNum = 1
	}
	return []int{minNum, maxNum}
}
