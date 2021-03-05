package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello","Alaska","Dad","Peace","omk"}))
}

func findWords(words []string) []string {
	keyboardArr := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	var res []string
	for _, word := range words {
		for i := range keyboardArr {
			if checkWordFromKeyboard(strings.ToLower(word), keyboardArr[i]) {
				res = append(res, word)
				break
			}
		}
	}
	return res
}

func checkWordFromKeyboard(word, keyboard string) bool {
	wordArr := []byte(word)
	for i := 0; i < len(wordArr); i++ {
		if !strings.Contains(string(keyboard), string(wordArr[i])) {
			return false
		}
	}
	return true
}
