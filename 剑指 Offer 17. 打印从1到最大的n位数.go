package main

import (
	"strconv"
	"strings"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 11:45 AM
**/

// 所有n位数字，可以看作0～9的全排列，只是1位时没有0、超过1位时不能有先导0，而且各位数字可以重复。只需要掌握全排列即可。

func printNumbers(n int) []int {
	var res []int
	arr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 1; i <= n; i++ {
		printNumbersNextPermutation(i, arr, "", &res)
	}
	return res
}

func printNumbersNextPermutation(n int, arr []string, s string, res *[]int) {
	if len(s) == n {
		t, _ := strconv.Atoi(s)
		*res = append(*res, t)
		return
	}
	for i := 0; i < len(arr); i++ {
		if len(s) == 0 && i == 0 {
			continue
		}
		printNumbersNextPermutation(n, arr, s+arr[i], res)
		strings.TrimSuffix(s, arr[i])
	}
}
