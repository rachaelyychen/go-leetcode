package main

// dp[i][j]表示到达下标（i，j）位置时，可以获得的最大值。动态方程：dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + grid[i][j]

func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
			}
			var left, up int
			if i > 0 {
				up = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			if up > left {
				dp[i][j] = up
			} else {
				dp[i][j] = left
			}
			dp[i][j] += grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

// func maxValue(grid [][]int) int {
// 	row, col, res := len(grid), len(grid[0]), 0
// 	var dfs func(i, j, cnt int)
// 	dfs = func(i, j, cnt int) {
// 		if i == row-1 && j == col-1 {
// 			if cnt+grid[i][j] > res {
// 				res = cnt + grid[i][j]
// 			}
// 			return
// 		}
// 		if i == row || j == col {
// 			return
// 		}
// 		dfs(i+1, j, cnt+grid[i][j])
// 		dfs(i, j+1, cnt+grid[i][j])
// 	}
// 	dfs(0, 0, 0)
// 	return res
// }
