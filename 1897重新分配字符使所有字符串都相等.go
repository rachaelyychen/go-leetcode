package main

func makeEqual(words []string) bool {
	arr := make([]int, 30)
	for i := range words {
		t := []byte(words[i])
		for _, b := range t {
			arr[b-'a'] += 1
		}
	}
	for i := range arr {
		if arr[i] > 0 && arr[i]%len(words) != 0 {
			return false
		}
	}
	return true
}
