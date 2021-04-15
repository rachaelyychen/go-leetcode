package main

func countConsistentStrings(allowed string, words []string) int {
	var res int
	m := make(map[byte]bool, 0)
	for _, b := range []byte(allowed) {
		m[b] = true
	}
	for i := range words {
		ok := true
		for _, b := range []byte(words[i]) {
			if _, exists := m[b]; !exists {
				ok = false
				break
			}
		}
		if ok {
			res += 1
		}
	}
	return res
}
