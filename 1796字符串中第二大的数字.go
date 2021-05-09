package main

func secondHighest(s string) int {
	arr := make([]int, 10)
	for _, b := range []byte(s) {
		if b >= '0' && b <= '9' {
			arr[b-'0'] = 1
		}
	}
	res, first := -1, false
	for i := 9; i >= 0; i-- {
		if arr[i] > 0 {
			if first {
				res = i
				break
			} else {
				first = true
			}
		}
	}
	return res
}
