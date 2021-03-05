package main

func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := low + (high - low) / 2
		t := num/mid
		if t == mid {
			if num % mid == 0 {
				return true
			}
			low = mid + 1
		} else if t < mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
