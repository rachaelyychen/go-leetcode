package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

func longestPalindrome(s string) int {
	charCntMap := make(map[byte]int, 0)
	arr := []byte(s)
	for i := range arr {
		if cnt, exists := charCntMap[arr[i]]; !exists {
			charCntMap[arr[i]] = 1
		} else {
			charCntMap[arr[i]] = cnt + 1
		}
	}
	singleCnt, doubleCnt := 0, 0
	for _, cnt := range charCntMap {
		if cnt%2 == 0 {
			doubleCnt += cnt
		} else {
			if singleCnt == 0 {
				singleCnt = 1
			}
			doubleCnt += cnt - 1
		}
	}
	return doubleCnt + singleCnt
}
