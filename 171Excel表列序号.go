package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	arr := []byte(s)
	var res int
	for i := range arr {
		res = res*26 + int(arr[i] - 'A') + 1
	}
	return res
}
