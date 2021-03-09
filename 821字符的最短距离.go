package main

func shortestToChar(s string, c byte) []int {
	var res []int
	arr := []byte(s)
	for i := range arr {
		var minDist int
		if arr[i] != c {
			before, after := -1, -1
			for j := i-1; j >= 0; j-- {
				if arr[j] == c {
					before = j
					break
				}
			}
			for j := i+1; j < len(arr); j++ {
				if arr[j] == c {
					after = j
					break
				}
			}
			if before == -1 {
				minDist = after - i
			} else if after == -1 {
				minDist = i - before
			} else {
				if i - before < after - i {
					minDist = i - before
				} else {
					minDist = after - i
				}
			}
		}
		res = append(res, minDist)
	}
	return res
}


