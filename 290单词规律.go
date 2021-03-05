package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	arrP := []byte(pattern)
	arrS := strings.Split(s, " ")
	byteStringMap := make(map[byte]string, 0)
	stringByteMap := make(map[string]byte, 0)
	if len(arrP) != len(arrS) {
		return false
	}
	for i := range arrP {
		if str, exists1 := byteStringMap[arrP[i]]; !exists1 {
			if _, exists2 := stringByteMap[arrS[i]]; !exists2 {
				byteStringMap[arrP[i]] = arrS[i]
				stringByteMap[arrS[i]] = arrP[i]
			} else {
				return false
			}
		} else {
			if arrS[i] != str {
				return false
			}
		}
	}
	return true
}
