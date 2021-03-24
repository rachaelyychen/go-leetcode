package main

func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if num == 0 {
			continue
		}
		var cnt int
		for num != 0 {
			num /= 10
			cnt += 1
		}
		if cnt%2 == 0 {
			res += 1
		}
	}
	return res
}


