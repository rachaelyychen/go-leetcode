package main

import (
	. "go-leetcode/kit"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	searchLeaf(root, "", &res)
	return res
}

func searchLeaf(root *TreeNode, s string, res *[]string) {
	if root == nil {
		return
	}
	s += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, s)
	} else {
		searchLeaf(root.Left, s+"->", res)
		searchLeaf(root.Right, s+"->", res)
	}
}
