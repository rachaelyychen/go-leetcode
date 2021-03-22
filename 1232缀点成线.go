package main

import "sort"

func checkStraightLine(coordinates [][]int) bool {
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][0] != coordinates[j][0] {
			return coordinates[i][0] < coordinates[j][0]
		}
		if coordinates[i][1] != coordinates[j][1] {
			return coordinates[i][1] < coordinates[j][1]
		}
		return true
	})
	for i := 0; i < len(coordinates); i++ {
		if i+2 < len(coordinates) && !isInLine(coordinates[i][0], coordinates[i][1],
			coordinates[i+1][0], coordinates[i+1][1], coordinates[i+2][0], coordinates[i+2][1]) {
			return false
		}
	}
	return true
}

func isInLine(x1, y1, x2, y2, x3, y3 int) bool {
	return (y2-y1)*(x3-x2) == (y3-y2)*(x2-x1)
}
