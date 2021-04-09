package main

func modifyString(s string) string {
	arr := []byte(s)
	for i := range arr {
		if arr[i] == '?' {
			for j := 0; j < 26; j++ {
				t := byte('a' + j)
				if (i-1 < 0 || i-1 >= 0 && arr[i-1] != t) && (i+1 >= len(arr) || i+1 < len(arr) && arr[i+1] != t) {
					arr[i] = t
					break
				}
			}
		}
	}
	return string(arr)
}
