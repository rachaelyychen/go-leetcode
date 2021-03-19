package main

func minCostToMoveChips(position []int) int {
	a, b := 0, 0
	for i := range position {
		if position[i] % 2 == 0 {
			a += 1
		} else {
			b += 1
		}
	}
	if a > b {
		return b
	}
	return a
}
