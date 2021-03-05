package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i - 1][j] == 1 {
					res -= 2
				}
				if j > 0 && grid[i][j - 1] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}
