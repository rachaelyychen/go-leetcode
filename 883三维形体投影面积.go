package main

func projectionArea(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		var colMax int
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				res += 1
			}
			if grid[i][j] > colMax {
				colMax = grid[i][j]
			}
		}
		res += colMax
	}
	for i := 0; i < len(grid[0]); i++ {
		var rowMax int
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > rowMax {
				rowMax = grid[j][i]
			}
		}
		res += rowMax
	}
	return res
}


