package main

func maximumWealth(accounts [][]int) int {
	var res int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
