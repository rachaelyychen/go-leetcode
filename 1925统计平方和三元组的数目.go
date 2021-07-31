package main

func countTriples(n int) int {
	var res int
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == j {
				continue
			}
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					res += 1
				}
			}
		}
	}
	return res
}
