package main

func firstBadVersion(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high - low) / 2
		if isBadVersion(mid) {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return high
}


