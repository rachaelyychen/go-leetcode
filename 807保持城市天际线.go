package main

// 遍历数组找到每行每列的最大值,对于每个元素,可增加到该行该列最大值中较小的一个

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	rowMaxArr, colMaxArr := make([]int, row), make([]int, col)
	for i := 0; i < row; i++ {
		var max int
		for j := 0; j < col; j++ {
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
		rowMaxArr[i] = max
	}
	for i := 0; i < col; i++ {
		var max int
		for j := 0; j < row; j++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		colMaxArr[i] = max
	}
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			t := rowMaxArr[i]
			if t > colMaxArr[j] {
				t = colMaxArr[j]
			}
			res += t - grid[i][j]
		}
	}
	return res
}
