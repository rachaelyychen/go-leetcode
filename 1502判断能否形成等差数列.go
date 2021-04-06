package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var t int
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			if i == 0 {
				t = arr[i+1] - arr[i]
			} else {
				if t != arr[i+1] - arr[i] {
					return false
				}
			}
		}
	}
	return true
}


