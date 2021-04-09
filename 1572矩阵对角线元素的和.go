package main

func diagonalSum(mat [][]int) int {
	var res int
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		t := len(mat) - 1 - i
		if t != i {
			res += mat[i][t]
		}
	}
	return res
}
