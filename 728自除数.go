package main

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		j := i
		ok := true
		for j != 0 {
			k := j%10
			if k == 0 || i % k != 0 {
				ok = false
				break
			}
			j /= 10
		}
		if ok {
			res = append(res, i)
		}
	}
	return res
}


