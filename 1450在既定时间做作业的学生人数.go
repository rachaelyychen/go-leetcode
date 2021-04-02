package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var res int
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res += 1
		}
	}
	return res
}


