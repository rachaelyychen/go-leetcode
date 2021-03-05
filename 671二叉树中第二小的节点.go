package main

import (
	. "go-leetcode/kit"
)

func findSecondMinimumValue(root *TreeNode) int {
	// root.val = min(root.left.val, root.right.val)
	var res int
	inOrder671(root, root.Val, &res)
	if res == 0 {
		res = -1
	}
	return res
}

func inOrder671(root *TreeNode, minVal int, res *int) {
	if root.Val > minVal {
		if *res == 0 || *res > root.Val {
			*res = root.Val
		}
	}
	if root.Left != nil {
		inOrder671(root.Left, minVal, res)
	}
	if root.Right != nil {
		inOrder671(root.Right, minVal, res)
	}
}
