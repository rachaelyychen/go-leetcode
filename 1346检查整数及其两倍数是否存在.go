package main

import "sort"

func checkIfExist(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr); i++ {
		if arr[len(arr)-1] < 2*arr[i] {
			break
		}
		for j := 0; j < len(arr); j++ {
			if j != i && arr[i]*2 == arr[j] {
				return true
			}
		}
	}
	return false
}
