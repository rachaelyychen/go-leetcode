package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aba"))
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteArr := []byte(ransomNote)
	magazineArr := []byte(magazine)
	magazineMap := make(map[byte]int, 0)
	for i := range magazineArr {
		if cnt, exists := magazineMap[magazineArr[i]]; !exists {
			magazineMap[magazineArr[i]] = 1
		} else {
			magazineMap[magazineArr[i]] = cnt + 1
		}
	}
	for i := range ransomNoteArr {
		if cnt, exists := magazineMap[ransomNoteArr[i]]; !exists {
			return false
		} else {
			if cnt == 0 {
				return false
			}
			magazineMap[ransomNoteArr[i]] = cnt - 1
		}
	}
	return true
}
