package main

import (
	"sort"
)

// 全排列。在遍历字符串的时候去重，保证一个位置上相同的字符只会被填入一次。
// 首先对原字符串排序，保证相同的字符都相邻，在递归函数中限制每次填入的字符一定是这个字符所在重复字符集合中「从左往右第一个未被填入的字符」。

func permutation(s string) []string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var res []string
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && arr[j-1] == arr[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, arr[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return res
}
