package main

func numSpecial(mat [][]int) int {
	var res int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				ok := true
				for k := 0; k < len(mat[i]); k++ {
					if k != j && mat[i][k] != 0 {
						ok = false
						break
					}
				}
				if ok == false {
					continue
				}
				for k := 0; k < len(mat); k++ {
					if k != i && mat[k][j] != 0 {
						ok = false
						break
					}
				}
				if ok == true {
					res += 1
				}
			}
		}
	}
	return res
}
