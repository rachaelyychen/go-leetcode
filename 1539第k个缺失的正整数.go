package main

func findKthPositive(arr []int, k int) int {
	var cnt int
	var j int
	for i := 1; i <= 3000; i++ {
		if j >= len(arr) || arr[j] != i {
			cnt += 1
			if cnt == k {
				return i
			}
		} else {
			j += 1
		}
	}
	return 0
}


