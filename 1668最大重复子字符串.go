package main

import "strings"

func maxRepeating(sequence string, word string) int {
	k := len(sequence) / len(word)
	for k > 0 {
		if strings.Contains(sequence, strings.Repeat(word, k)) {
			return k
		}
		k -= 1
	}
	return 0
}
