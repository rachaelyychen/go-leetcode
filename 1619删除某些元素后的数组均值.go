package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	cnt, sum := len(arr)/20, 0
	for i := cnt; i < len(arr)-cnt; i++ {
		sum += arr[i]
	}
	return 1.0 * float64(sum) / (0.9 * float64(len(arr)))
}
