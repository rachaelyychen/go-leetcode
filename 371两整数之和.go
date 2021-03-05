package main

func getSum(a int, b int) int {
	for b != 0 {
		c := int((uint(a) & uint(b)) << 1)
		a ^= b
		b = c
	}
	return a
}
