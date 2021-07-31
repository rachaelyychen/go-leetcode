package main

func lengthOfLongestSubstring(s string) int {
	arr := []byte(s)
	var res int
	for i := 0; i < len(arr); i++ {
		m := make(map[byte]int, 0)
		m[arr[i]] = 1
		for j := i + 1; j <= len(arr); j++ {
			if j == len(arr) {
				if res < len(m) {
					res = len(m)
				}
				break
			}
			if _, exists := m[arr[j]]; exists {
				if res < len(m) {
					res = len(m)
				}
				break
			} else {
				m[arr[j]] = 1
			}
		}
	}
	return res
}
