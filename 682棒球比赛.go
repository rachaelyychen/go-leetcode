package main

import "strconv"

func calPoints(ops []string) int {
	var arr []int
	for i := range ops {
		switch ops[i] {
		case "+":
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
		case "D":
			arr = append(arr, arr[len(arr)-1]*2)
		case "C":
			arr = arr[:len(arr)-1]
		default:
			x, _ := strconv.Atoi(ops[i])
			arr = append(arr, x)
		}
	}
	var res int
	for i := range arr {
		res += arr[i]
	}
	return res
}
