package main

func oddCells(m int, n int, indices [][]int) int {
	arrRow, arrCol := make([]int, 50), make([]int, 50)
	for i := range indices {
		arrRow[indices[i][0]] += 1
		arrCol[indices[i][1]] += 1
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (arrRow[i]+arrCol[j]) % 2 != 0 {
				res += 1
			}
		}
	}
	return res
}


