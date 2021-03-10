package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return isRectangle(rec1) && isRectangle(rec2) &&
		!(rec1[2] <= rec2[0] || rec2[2] <= rec1[0] || rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
}

func isRectangle(rec []int) bool {
	return !(rec[0] == rec[2] || rec[1] == rec[3])
}
