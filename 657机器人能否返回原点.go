package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeCircle("UDR"))
}

func judgeCircle(moves string) bool {
	arr := []byte(moves)
	moveCntMap := make(map[byte]int, 0)
	for i := range arr {
		switch arr[i] {
		case 'U':
			if cnt, exists := moveCntMap['U']; !exists {
				moveCntMap['U'] = 1
			} else {
				moveCntMap['U'] = cnt + 1
			}
		case 'D':
			if cnt, exists := moveCntMap['U']; !exists {
				moveCntMap['U'] = -1
			} else {
				moveCntMap['U'] = cnt - 1
			}
		case 'L':
			if cnt, exists := moveCntMap['L']; !exists {
				moveCntMap['L'] = 1
			} else {
				moveCntMap['L'] = cnt + 1
			}
		case 'R':
			if cnt, exists := moveCntMap['L']; !exists {
				moveCntMap['L'] = -1
			} else {
				moveCntMap['L'] = cnt - 1
			}
		}
	}
	for _, cnt := range moveCntMap {
		if cnt != 0 {
			return false
		}
	}
	return true
}
