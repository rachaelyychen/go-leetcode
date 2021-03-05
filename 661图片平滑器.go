package main

func imageSmoother(M [][]int) [][]int {
	row := len(M)
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		col := len(M[i])
		res[i] = make([]int, col)
		for j := 0; j < col; j++ {
			sum, cnt := M[i][j], 1
			if i-1 >= 0 && j-1 >= 0 {
				sum += M[i-1][j-1]
				cnt += 1
			}
			if i-1 >= 0 {
				sum += M[i-1][j]
				cnt += 1
			}
			if i-1 >= 0 && j+1 < col {
				sum += M[i-1][j+1]
				cnt += 1
			}
			if j-1 >= 0 {
				sum += M[i][j-1]
				cnt += 1
			}
			if j+1 < col {
				sum += M[i][j+1]
				cnt += 1
			}
			if i+1 < row && j-1 >= 0 {
				sum += M[i+1][j-1]
				cnt += 1
			}
			if i+1 < row {
				sum += M[i+1][j]
				cnt += 1
			}
			if i+1 < row && j+1 < col {
				sum += M[i+1][j+1]
				cnt += 1
			}
			res[i][j] = sum / cnt
		}
	}
	return res
}
