package main

func canBeEqual(target []int, arr []int) bool {
	t := make([]int, 1005)
	for i := range target {
		t[target[i]] += 1
	}
	for i := range arr {
		t[arr[i]] -= 1
	}
	for i := range t {
		if t[i] != 0 {
			return false
		}
	}
	return true
}


