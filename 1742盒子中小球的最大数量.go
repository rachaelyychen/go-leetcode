package main

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int, 0)
	var maxCnt int
	for i := lowLimit; i <= highLimit; i++ {
		t, sum := i, 0
		for t > 0 {
			sum += t % 10
			t /= 10
		}
		if cnt, exists := m[sum]; !exists {
			m[sum] = 1
			if 1 > maxCnt {
				maxCnt = 1
			}
		} else {
			m[sum] = cnt + 1
			if cnt+1 > maxCnt {
				maxCnt = cnt + 1
			}
		}
	}
	return maxCnt
}
