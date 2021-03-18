package main

func countCharacters(words []string, chars string) int {
	arr := make([]int, 30)
	for _, b := range []byte(chars) {
		arr[b-'a'] += 1
	}
	var res int
	for _, word := range words {
		t := make([]int, 30)
		for _, b := range []byte(word) {
			t[b-'a'] += 1
		}
		ok := true
		for i := range t {
			if t[i] > arr[i] {
				ok = false
				break
			}
		}
		if ok {
			res += len(word)
		}
	}
	return res
}


