package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("cc"))
}

func firstUniqChar(s string) int {
	arr := []byte(s)
	charIndexMap := make(map[byte]int, 0)
	for i := range arr {
		if _, exists := charIndexMap[arr[i]]; !exists {
			charIndexMap[arr[i]] = i
		} else {
			charIndexMap[arr[i]] = -1
		}
	}
	res := len(arr)
	for _, index := range charIndexMap {
		if index > -1 && index < res{
			res = index
		}
	}
	if res == len(arr) {
		res = -1
	}
	return res
}
