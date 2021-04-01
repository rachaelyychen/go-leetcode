package main

func maxPower(s string) int {
	var res int
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		cnt := 1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] != arr[i] {
				i = j - 1
				break
			}
			cnt += 1
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
