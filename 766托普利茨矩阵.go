package main

func isToeplitzMatrix(matrix [][]int) bool {
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				if matrix[i-1][j-1] != matrix[i][j] {
					return false
				}
			}
		}
	}
	return true
}


