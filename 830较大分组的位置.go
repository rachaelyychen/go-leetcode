package main

func largeGroupPositions(s string) [][]int {
	var res [][]int
	arr := []byte(s)
	start, end := 0, 1
	for end <= len(arr) {
		if end == len(arr) || arr[end] != arr[start] {
			if end-start >= 3 {
				res = append(res, []int{start, end-1})
				start = end
			}
		}
		end += 1
	}
	return res
}
