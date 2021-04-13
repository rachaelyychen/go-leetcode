package main

func maxLengthBetweenEqualCharacters(s string) int {
	maxLength := -1
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				t := j-i-1
				if t > maxLength {
					maxLength = t
				}
			}
		}
	}
	return maxLength
}


