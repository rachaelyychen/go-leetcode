package main

import "fmt"

func main() {
	fmt.Println(maxCount(4e4, 4e4, [][]int{{2, 2}, {3, 3}}))
}

func maxCount(m int, n int, ops [][]int) int {
	minA, minB := 0, 0
	for i := range ops {
		if len(ops[i]) == 2 {
			if minA == 0 || minA > ops[i][0] {
				minA = ops[i][0]
			}
			if minB == 0 || minB > ops[i][1] {
				minB = ops[i][1]
			}
		}
	}
	if minA*minB == 0 {
		return m * n
	}
	return minA * minB
}
