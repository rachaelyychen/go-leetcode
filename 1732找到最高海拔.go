package main

func largestAltitude(gain []int) int {
	arr := make([]int, len(gain))
	for i := range gain {
		if i == 0 {
			arr[i] = gain[i]
		} else {
			arr[i] = gain[i] + arr[i-1]
		}
	}
	var res int
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
