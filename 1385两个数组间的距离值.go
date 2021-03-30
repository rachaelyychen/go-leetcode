package main

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	for i := range arr1 {
		var ok bool
		for j := range arr2 {
			var t int
			if arr1[i] > arr2[j] {
				t = arr1[i] - arr2[j]
			} else {
				t = arr2[j] - arr1[i]
			}
			if t <= d {
				ok = true
				break
			}
		}
		if !ok {
			res += 1
		}
	}
	return res
}
