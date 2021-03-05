package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	vowelMap := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	arr := []byte(s)
	i, j := 0, len(arr)-1
	for i < j {
		_, e1 := vowelMap[arr[i]]
		_, e2 := vowelMap[arr[j]]
		if e1 && e2 {
			t := arr[i]
			arr[i] = arr[j]
			arr[j] = t
			i += 1
			j -= 1
		} else if !e1 && e2 {
			i += 1
		} else if e1 && !e2 {
			j -= 1
		} else {
			i += 1
			j -= 1
		}
	}
	return string(arr)
}
