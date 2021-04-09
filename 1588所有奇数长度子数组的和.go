package main

func sumOddLengthSubarrays(arr []int) int {
	var res int
	for l := 1; l <= len(arr); l += 2 {
		for i := 0; i < len(arr); i++ {
			if i+l <= len(arr) {
				res += sumSubarray(arr, i, i+l)
			}
		}
	}
	return res
}

func sumSubarray(arr []int, from, to int) int {
	var res int
	for i := from; i < to; i++ {
		res += arr[i]
	}
	return res
}
