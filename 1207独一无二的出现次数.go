package main

func uniqueOccurrences(arr []int) bool {
	arr1, arr2 := make([]int, 2001), make([]int, 1001)
	for i := range arr {
		arr1[arr[i]+1000] += 1
	}
	for i := range arr1 {
		if arr1[i] > 0 {
			arr2[arr1[i]] += 1
		}
	}
	for i := range arr2 {
		if arr2[i] > 1 {
			return false
		}
	}
	return true
}


