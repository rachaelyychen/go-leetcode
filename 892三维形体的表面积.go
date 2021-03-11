package main

func surfaceArea(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				res += 6 * grid[i][j]
				if grid[i][j] > 1 {
					res -= 2 * (grid[i][j] - 1)
				}
			}
			if j+1 < len(grid[i]) && grid[i][j+1] > 0 {
				if grid[i][j] < grid[i][j+1] {
					res -= 2 * (grid[i][j])
				} else {
					res -= 2 * (grid[i][j+1])
				}
			}
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if j+1 < len(grid) && grid[j+1][i] > 0 {
				if grid[j][i] < grid[j+1][i] {
					res -= 2 * (grid[j][i])
				} else {
					res -= 2 * (grid[j+1][i])
				}
			}
		}
	}
	return res
}
