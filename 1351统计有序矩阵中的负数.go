package main

func countNegatives(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i])-1; j >= 0; j-- {
			if grid[i][j] < 0 {
				res += 1
			} else {
				break
			}
		}
	}
	return res
}


