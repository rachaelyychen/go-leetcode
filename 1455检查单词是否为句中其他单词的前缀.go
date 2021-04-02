package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	res := -1
	arr := strings.Split(sentence, " ")
	for i := range arr {
		if strings.HasPrefix(arr[i], searchWord) {
			res = i+1
			break
		}
	}
	return res
}


