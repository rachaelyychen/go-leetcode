package main

func countGoodRectangles(rectangles [][]int) int {
	res, maxLen, arr := 0, 0, make([]int, len(rectangles))
	for i := range rectangles {
		var t int
		if rectangles[i][0] > rectangles[i][1] {
			t = rectangles[i][1]
		} else {
			t = rectangles[i][0]
		}
		arr[i] = t
		if t > maxLen {
			maxLen = t
		}
	}
	for i := range arr {
		if arr[i] == maxLen {
			res += 1
		}
	}
	return res
}
