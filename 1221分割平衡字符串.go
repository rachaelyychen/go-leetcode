package main

func balancedStringSplit(s string) int {
	num, res := 0, 0
	for _, b := range []byte(s) {
		switch b {
		case 'R':
			num += 1
		case 'L':
			num -= 1
		}
		if num == 0 {
			res += 1
		}
	}
	return res
}
