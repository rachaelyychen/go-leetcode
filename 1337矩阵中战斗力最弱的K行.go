package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	type WeakRow struct {
		weak int
		row int
	}
	var arr []WeakRow
	for i := 0; i < len(mat); i++ {
		var weak int
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				weak += 1
			} else {
				break
			}
		}
		arr = append(arr, WeakRow{weak:weak, row:i})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].weak != arr[j].weak {
			return arr[i].weak < arr[j].weak
		}
		return arr[i].row < arr[j].row
	})
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, arr[i].row)
	}
	return res
}


