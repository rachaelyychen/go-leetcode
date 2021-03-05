package main

import "unicode"

func detectCapitalUse(word string) bool {
	arr := []rune(word)
	// 全部大写
	// 全部小写
	// 至少一个字母时只有首字母大写
	return isAllUpper(arr) || isAllLower(arr) || isFirstUpper(arr)
}

func isAllUpper(arr []rune) bool {
	for i := range arr {
		if !unicode.IsUpper(arr[i]) {
			return false
		}
	}
	return true
}

func isAllLower(arr []rune) bool {
	for i := range arr {
		if !unicode.IsLower(arr[i]) {
			return false
		}
	}
	return true
}

func isFirstUpper(arr []rune) bool {
	for i := range arr {
		if i == 0 && !unicode.IsUpper(arr[i]) {
			return false
		}
		if i > 0 && !unicode.IsLower(arr[i]) {
			return false
		}
	}
	return true
}
