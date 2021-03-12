package main

func diStringMatch(S string) []int {
	i, j := 0, len(S)
	var res []int
	for _, b := range []byte(S) {
		if b == 'I' {
			res = append(res, i)
			i += 1
		}
		if b == 'D' {
			res = append(res, j)
			j -= 1
		}
	}
	res = append(res, j)
	return res
}
