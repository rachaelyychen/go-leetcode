package main

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	var cnt int
	words := strings.Split(text, " ")
	for i := range words {
		arr := []rune(words[i])
		for i := range arr {
			if strings.ContainsRune(brokenLetters, arr[i]) {
				cnt += 1
				break
			}
		}
	}
	return len(words) - cnt
}
