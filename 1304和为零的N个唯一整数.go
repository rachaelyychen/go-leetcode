package main

func sumZero(n int) []int {
	var res []int
	var sum int
	for i := 1; i < n; i++ {
		res = append(res, i)
		sum += i
	}
	res = append(res, -sum)
	return res
}
