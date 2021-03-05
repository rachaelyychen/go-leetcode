package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(findTheDifference("", "a")))
}

func findTheDifference(s string, t string) byte {
	sArr := []byte(s)
	tArr := []byte(t)
	charCntMap := make(map[byte]int, len(sArr))
	for i := range sArr {
		if cnt, exists := charCntMap[sArr[i]]; !exists {
			charCntMap[sArr[i]] = 1
		} else {
			charCntMap[sArr[i]] = cnt + 1
		}
	}
	for i := range tArr {
		if cnt, exists := charCntMap[tArr[i]]; !exists {
			return tArr[i]
		} else {
			cnt -= 1
			if cnt < 0 {
				return tArr[i]
			}
			charCntMap[tArr[i]] = cnt
		}
	}
	return ' '
}
