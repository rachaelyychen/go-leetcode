package main

func firstUniqChar(s string) byte {
	arr := make([]int, 26)
	for i := range s {
		arr[s[i]-'a'] += 1
	}
	var res byte = ' '
	for i := range s {
		if arr[s[i]-'a'] == 1 {
			res = s[i]
			break
		}
	}
	return res
}
