package main

func minTimeToVisitAllPoints(points [][]int) int {
	var res int
	for i := 0; i < len(points); i++ {
		if i+1 < len(points) {
			res += minTimeForPoints(points[i], points[i+1])
		}
	}
	return res
}

func minTimeForPoints(x, y []int) int {
	var res int
	for x[0] != y[0] || x[1] != y[1] {
		if x[0] == y[0] {
			if x[1] > y[1] {
				res += x[1] - y[1]
			} else {
				res += y[1] - x[1]
			}
			x[1] = y[1]
			continue
		}
		if x[1] == y[1] {
			if x[0] > y[0] {
				res += x[0] - y[0]
			} else {
				res += y[0] - x[0]
			}
			x[0] = y[0]
			continue
		}
		if x[0] > y[0] {
			x[0] -= 1
		} else {
			x[0] += 1
		}
		if x[1] > y[1] {
			x[1] -= 1
		} else {
			x[1] += 1
		}
		res += 1
	}
	return res
}


