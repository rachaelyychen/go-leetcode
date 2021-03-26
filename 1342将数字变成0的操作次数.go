package main

func numberOfSteps (num int) int {
	var res int
	for num != 0 {
		if num%2 ==0 {
			num /= 2
		} else {
			num -= 1
		}
		res += 1
	}
	return res
}


