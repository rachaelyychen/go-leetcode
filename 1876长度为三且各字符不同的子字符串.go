package main

func countGoodSubstrings(s string) int {
	arr, res := []byte(s), 0
	for i := 0; i < len(arr); i++ {
		if i+2< len(arr) && arr[i] != arr[i+1] && arr[i] != arr[i+2] && arr[i+1] != arr[i+2] {
			res += 1
		}
	}
	return res
}
