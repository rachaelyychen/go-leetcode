package main

import "strconv"

var digitLetter = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	var res []string
	dfs([]byte(digits), 0, "", &res)
	return res
}

func dfs(arr []byte, index int, t string, res *[]string) {
	if index == len(arr) {
		*res = append(*res, t)
		return
	}
	n, _ := strconv.Atoi(string(arr[index]))
	arr1 := []byte(digitLetter[n])
	for i := range arr1 {
		dfs(arr, index+1, t+string(arr1[i]), res)
	}
}
