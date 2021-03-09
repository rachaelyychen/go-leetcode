package main

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ReplaceAll(paragraph, "!", " ")
	paragraph = strings.ReplaceAll(paragraph, "?", " ")
	paragraph = strings.ReplaceAll(paragraph, "'", " ")
	paragraph = strings.ReplaceAll(paragraph, ",", " ")
	paragraph = strings.ReplaceAll(paragraph, ";", " ")
	paragraph = strings.ReplaceAll(paragraph, ".", " ")
	paragraph = strings.ToLower(paragraph)
	words := strings.Split(paragraph, " ")
	wordCntMap := make(map[string]int, 0)
	for _, b := range banned {
		wordCntMap[b] = -1
	}
	var res string
	var maxCnt int
	for _, word := range words {
		if len(strings.TrimSpace(word)) == 0 {
			continue
		}
		if cnt, exists := wordCntMap[word]; !exists {
			wordCntMap[word] = 1
		} else {
			if cnt >= 0 {
				wordCntMap[word] = cnt + 1
			}
		}
		if wordCntMap[word] > maxCnt {
			maxCnt = wordCntMap[word]
			res = word
		}
	}
	return res
}
