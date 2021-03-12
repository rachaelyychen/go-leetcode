package main

func minDeletionSize(strs []string) int {
	row, col := len(strs), len(strs[0])
	var res int
	for i := 0; i < col; i++ {
		for j := 0; j < row-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				res += 1
				break
			}
		}
	}
	return res
}
