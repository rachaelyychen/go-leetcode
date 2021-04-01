package main

func buildArray(target []int, n int) []string {
	var res []string
	i, j := 0, 1
	for {
		if i == len(target) || j > n {
			break
		}
		if target[i] != j {
			res = append(res, "Push", "Pop")
		} else {
			res = append(res, "Push")
			i += 1
		}
		j += 1
	}
	return res
}


