package main

func findRotation(mat [][]int, target [][]int) bool {
	var cnt int
	n := len(mat)
	for !isRotationEqual(mat, target, n) {
		if cnt == 3 {
			return false
		}
		cnt += 1
		var x [][]int
		for i := 0; i < n; i++ {
			t := make([]int, n)
			for j := 0; j < n; j++ {
				t[j] = mat[n-j-1][i]
			}
			x = append(x, t)
		}
		mat = x
	}
	return true
}

func isRotationEqual(mat [][]int, target [][]int, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}
