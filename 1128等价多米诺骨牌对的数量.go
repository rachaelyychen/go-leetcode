package main

import (
	"fmt"
)

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1,2},{2,1},{1,2},{2,1}}))
}

func numEquivDominoPairs(dominoes [][]int) int {
	a := make([]int, 100)
	count := 0
	for i := range dominoes {
		num1 := dominoes[i][0]*10+dominoes[i][1]
		num2 := dominoes[i][1]*10+dominoes[i][0]
		if num2 >= num1 {
			a[num1] += 1
		} else  {
			a[num2] += 1
		}
	}
	for i := range a {
		if a[i] > 0 {
			count += getRelations(a[i])
		}
	}
	return count
}

func getRelations(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		count += i
	}
	return count
}
