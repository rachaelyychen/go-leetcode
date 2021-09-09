package main

func countBits(n int) []int {
	var res []int
	res = append(res, 0)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			res = append(res, res[len(res)-1]+1)
		} else {
			res = append(res, res[i/2])
		}
	}
	return res
}
