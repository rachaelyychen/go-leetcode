package main

func maximumPopulation(logs [][]int) int {
	arr, max := make([]int, 105), 0
	for i := 1950; i <= 2050; i++ {
		var t int
		for j := range logs {
			if logs[j][0] <= i && i < logs[j][1] {
				t += 1
			}
		}
		arr[i-1950] = t
		if t > max {
			max = t
		}
	}
	for i := 1950; i <= 2050; i++ {
		if arr[i-1950] == max {
			return i
		}
	}
	return 0
}
