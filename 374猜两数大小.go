package main

func guessNumber(n int) int {
	low, high := 1, n
	for low <= high {
		mid := (low + high) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0
}
