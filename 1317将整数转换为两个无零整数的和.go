package main

func getNoZeroIntegers(n int) []int {
	i, j := 1, n-1
	for i <= j {
		if i+j == n && isZeroInteger(i) && isZeroInteger(j) {
			return []int{i, j}
		}
		i += 1
		j -= 1
	}
	return []int{}
}

func isZeroInteger(x int) bool {
	for x != 0 {
		if x%10 == 0 {
			return false
		}
		x /= 10
	}
	return true
}


