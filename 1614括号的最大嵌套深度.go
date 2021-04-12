package main

func maxDepth(s string) int {
	cnt, max := 0, 0
	for _, b := range []byte(s) {
		if b == '(' {
			cnt += 1
			if cnt > max {
				max = cnt
			}
		}
		if b == ')' {
			cnt -= 1
		}
	}
	return max
}


