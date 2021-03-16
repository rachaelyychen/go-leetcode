package main

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	type cellDist struct {
		r int
		c int
		d int
	}
	var arr []cellDist
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if i != r0 || j != c0 {
				arr = append(arr, cellDist{r: i, c: j, d: getCellsDist(i, j, r0, c0)})
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].d < arr[j].d
	})
	res := [][]int{{r0, c0}}
	for i := range arr {
		res = append(res, []int{arr[i].r, arr[i].c})
	}
	return res
}

func getCellsDist(r1, c1, r2, c2 int) int {
	var res int
	if r1 < r2 {
		res += r2 - r1
	} else {
		res += r1 - r2
	}
	if c1 < c2 {
		res += c2 - c1
	} else {
		res += c1 - c2
	}
	return res
}
