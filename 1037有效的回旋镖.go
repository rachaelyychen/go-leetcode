package main

import (
	"fmt"
)

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(isBoomerang([][]int{{0, 2}, {2, 1}, {0, 0}}))
}

func isBoomerang(points [][]int) bool {
	if isSamePoint(points[0], points[1]) || isSamePoint(points[0], points[2]) ||
		isSamePoint(points[1], points[2]) {
		return false
	}
	if isPointsInLine(points[0], points[1], points[2]) {
		return false
	}
	return true
}

func isSamePoint(x, y []int) bool {
	if x[0] == y[0] && x[1] == y[1] {
		return true
	}
	return false
}

func isPointsInLine(x, y, z []int) bool {
	if (x[0]-y[0])*(y[1]-z[1]) == (x[1]-y[1])*(y[0]-z[0]) {
		return true
	}
	return false
}
