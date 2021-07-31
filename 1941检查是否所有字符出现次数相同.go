package main

func areOccurrencesEqual(s string) bool {
	arr1, arr2 := make([]int, 30), []byte(s)
	var cnt int
	for i := range arr2 {
		b := arr2[i] - 'a'
		arr1[b] += 1
		if arr1[b] > cnt {
			cnt = arr1[b]
		}
	}
	for i := range arr1 {
		if arr1[i] > 0 && arr1[i] != cnt {
			return false
		}
	}
	return true
}
