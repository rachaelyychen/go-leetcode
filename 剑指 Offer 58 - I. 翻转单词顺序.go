package main

import "strings"

// 双指针确定单词首尾字符，逆序遍历字符串，跳过多余空格，将单词依次拼接返回。

func reverseWords(s string) string {
	var res string
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i -= 1
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i -= 1
		}
		res += s[i+1:j+1] + " "
	}
	return strings.TrimSpace(res)
}
