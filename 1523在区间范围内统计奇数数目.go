package main

func countOdds(low int, high int) int {
	var res int
	if low%2 == 1 || high%2 == 1 {
		res += 1
	}
	return res + (high-low)/2
}
