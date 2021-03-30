package main

func findLucky(arr []int) int {
	cntArr := make([]int, 501)
	for i := range arr {
		cntArr[arr[i]] += 1
	}
	for i := len(cntArr) - 1; i >= 0; i-- {
		if cntArr[i] > 0 && cntArr[i] == i {
			return i
		}
	}
	return -1
}
