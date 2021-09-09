package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	var res int
	lastEndTime := -1
	for i := range timeSeries {
		if lastEndTime == -1 {
			res += duration
		} else {
			if lastEndTime >= timeSeries[i] {
				res += duration - (lastEndTime - timeSeries[i] + 1)
			} else {
				res += duration
			}
		}
		lastEndTime = timeSeries[i] + duration - 1
	}
	return res
}
