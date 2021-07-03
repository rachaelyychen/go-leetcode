package main

import (
	"strconv"
	"strings"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return getNumFromWord(firstWord)+getNumFromWord(secondWord) == getNumFromWord(targetWord)
}

func getNumFromWord(word string) int {
	for strings.HasPrefix(word, "a") {
		word = strings.TrimPrefix(word, "a")
	}
	arr, res := []byte(word), ""
	for i := range arr {
		res += strconv.Itoa(int(arr[i] - 'a'))
	}
	num, _ := strconv.Atoi(res)
	return num
}
