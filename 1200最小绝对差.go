package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	minAbsDiff := arr[len(arr)-1] - arr[0]
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && arr[i+1]-arr[i] < minAbsDiff {
			minAbsDiff = arr[i+1]-arr[i]
		}
	}
	var res [][]int
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && arr[i+1]-arr[i] == minAbsDiff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}


