package main

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	for i := range code {
		var sum int
		if k == 0 {
			sum = 0
		} else if k > 0 {
			for j := i + 1; j <= i+k; j++ {
				sum += code[j%len(code)]
			}
		} else {
			for j := i - 1; j >= i+k; j-- {
				sum += code[(j+len(code))%len(code)]
			}
		}
		res[i] = sum
	}
	return res
}
