package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	x := 1
	for i := 0; i <= rowIndex; i++ {
		res[i] = x
		x = x* (rowIndex-i) / (i+1)
	}
	return res
}
