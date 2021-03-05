package main

import (
	. "go-leetcode/kit"
	"strconv"
	"strings"
)

func tree2str(t *TreeNode) string {
	var res string
	inOrder606(t, &res)
	return strings.TrimSuffix(res, ")")
}

func inOrder606(t *TreeNode, res *string) {
	if t == nil {
		return
	}
	*res += strconv.Itoa(t.Val)
	if t.Left != nil || t.Left == nil && t.Right != nil {
		*res += "("
		if t.Left != nil {
			inOrder606(t.Left, res)
		} else {
			*res += ")"
		}
	}
	if t.Right != nil {
		*res += "("
		inOrder606(t.Right, res)
	}
	*res += ")"
}
