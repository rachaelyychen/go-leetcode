package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var res [][]int
	for i := 1; i <= numRows; i++ {
		generateNumRows(i, &res)
	}
	return res
}

func generateNumRows(numRows int, res *[][]int) {
	row := make([]int, numRows)
	row[0] = 1
	row[numRows-1] = 1
	if numRows > 2 {
		for i := 1; i < numRows-1; i++ {
			row[i] = (*res)[numRows-2][i-1] + (*res)[numRows-2][i]
		}
	}
	*res = append(*res, row)
}
