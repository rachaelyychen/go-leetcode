package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	wordCntMap := make(map[string]int, 0)
	arrA, arrB := strings.Split(A, " "), strings.Split(B, " ")
	for _, word := range arrA {
		if _, exists := wordCntMap[word]; !exists {
			wordCntMap[word] = 1
		} else {
			wordCntMap[word] = 0
		}
	}
	for _, word := range arrB {
		if _, exists := wordCntMap[word]; !exists {
			wordCntMap[word] = -1
		} else {
			wordCntMap[word] = 0
		}
	}
	var res []string
	for word, cnt := range wordCntMap {
		if cnt != 0 {
			res = append(res, word)
		}
	}
	return res
}



