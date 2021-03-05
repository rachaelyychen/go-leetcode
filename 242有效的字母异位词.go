package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	runeCntMap := make(map[rune]int, 0)
	sr := []rune(s)
	for i := range sr {
		if cnt, exists := runeCntMap[sr[i]]; !exists {
			runeCntMap[sr[i]] = 1
		} else {
			runeCntMap[sr[i]] = cnt + 1
		}
	}
	st := []rune(t)
	for i := range st {
		if cnt, exists := runeCntMap[st[i]]; !exists {
			return false
		} else {
			runeCntMap[st[i]] = cnt - 1
		}
	}
	for _, v := range runeCntMap {
		if v != 0 {
			return false
		}
	}
	return true
}
