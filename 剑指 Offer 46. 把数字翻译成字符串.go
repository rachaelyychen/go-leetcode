package main

import "strconv"

func translateNum(num int) int {
	arr := []byte(strconv.Itoa(num))
	var res int
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(arr) {
			res += 1
			return
		}
		if index > len(arr) {
			return
		}
		dfs(index + 1)
		if arr[index] > '2' || arr[index] == '0' || arr[index] == '2' && index+1 < len(arr) && arr[index+1] > '5' {
			return
		}
		dfs(index + 2)
	}
	dfs(0)
	return res
}
