package main

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i < len(res); i++ {
		res[i] = encoded[i-1] ^ res[i-1]
	}
	return res
}
