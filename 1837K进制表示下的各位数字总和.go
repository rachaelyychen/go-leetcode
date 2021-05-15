package main

func sumBase(n int, k int) int {
	var res int
	for n > 0 {
		res += n % k
		n /= k
	}
	return res
}
