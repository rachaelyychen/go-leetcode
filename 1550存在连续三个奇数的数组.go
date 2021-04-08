package main

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if i+2 < len(arr) && isOdd(arr[i]) && isOdd(arr[i+1]) && isOdd(arr[i+2]) {
			return true
		}
	}
	return false
}

func isOdd(x int) bool {
	if x > 0 && x%2 > 0 {
		return true
	}
	return false
}


