package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	romanIntMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var res int
	arr := []rune(s)
	length := len(arr)
	for i := range arr {
		if i < length-1 && romanIntMap[string(arr[i])] < romanIntMap[string(arr[i+1])] {
			res -= romanIntMap[string(arr[i])]
		} else {
			res += romanIntMap[string(arr[i])]
		}
	}
	return res
}
