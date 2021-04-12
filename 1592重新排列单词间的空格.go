package main

import (
	"strings"
)

func reorderSpaces(text string) string {
	var spaceCnt int
	var words []string
	var word string
	var inWord bool
	arr := []byte(text)
	for i := range arr {
		if arr[i] == ' ' {
			spaceCnt += 1
			if inWord {
				inWord = false
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(arr[i])
			if !inWord {
				inWord = true
			}
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	var t int
	if len(words) == 1 {
		t = 0
	} else {
		t = spaceCnt / (len(words) - 1)
	}
	var res string
	for i := range words {
		if i > 0 {
			res += strings.Repeat(" ", t)
			spaceCnt -= t
		}
		res += words[i]
	}
	res += strings.Repeat(" ", spaceCnt)
	return res
}
