package main

func replaceElements(arr []int) []int {
	res, max := make([]int, len(arr)), -1
	res[len(arr)-1] = max
	for i := len(arr)-2; i >= 0; i-- {
		if max < arr[i+1] {
			max = arr[i+1]
		}
		res[i] = max
	}
	return res
}


