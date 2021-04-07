package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	len := len(arr)
	var res int
	for i := 0; i <= len-3; i++ {
		for j := i + 1; j <= len-2; j++ {
			if abs(arr[i], arr[j]) <= a {
				for k := j + 1; k <= len-1; k++ {
					if abs(arr[j], arr[k]) <= b && abs(arr[i], arr[k]) <= c {
						res++
					}
				}
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
