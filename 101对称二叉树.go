package main

import (
	. "go-leetcode/kit"
)

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTree(root, root)
}

func isSymmetricTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return isSymmetricTree(t1.Left, t2.Right) && isSymmetricTree(t1.Right, t2.Left)
}
