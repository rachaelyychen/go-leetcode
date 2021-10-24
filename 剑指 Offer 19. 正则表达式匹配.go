package main

import "regexp"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 9:57 PM
**/

// regexp.Match()方法是检查s是否包含能被p匹配的子串，对于s="aa"而p="a"来说，返回true。

func isMatch(s string, p string) bool {
	re, _ := regexp.Compile(p)
	loc := re.FindStringIndex(s)
	return len(loc) > 0 && loc[0] == 0 && loc[1] == len(s)
}
