package main

import . "go-leetcode/kit"

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return preOrder965(root, root.Val)
}

func preOrder965(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val != val {
		return false
	}
	ok := true
	if root.Left != nil {
		ok = preOrder965(root.Left, val)
	}
	if !ok {
		return false
	}
	if root.Right != nil {
		ok = preOrder965(root.Right, val)
	}
	return ok
}


