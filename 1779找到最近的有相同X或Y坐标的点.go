package main

func nearestValidPoint(x int, y int, points [][]int) int {
	res, minDist := -1, -1
	for i := range points {
		if points[i][0] == x || points[i][1] == y {
			var d int
			if points[i][0] > x {
				d = points[i][0] - x
			} else {
				d = x - points[i][0]
			}
			if points[i][1] > y {
				d += points[i][1] - y
			} else {
				d += y - points[i][1]
			}
			if minDist < 0 || minDist > d {
				res = i
				minDist = d
			}
		}
	}
	return res
}
