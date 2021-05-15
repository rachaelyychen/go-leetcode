package main

func checkIfPangram(sentence string) bool {
	t, arr := make([]int, 26), []byte(sentence)
	for i := range arr {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			t[arr[i]-'a'] += 1
		}
	}
	for i := range t {
		if t[i] == 0 {
			return false
		}
	}
	return true
}
