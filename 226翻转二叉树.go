package main

import . "go-leetcode/kit"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val: root.Val,
		Left: invertTree(root.Right),
		Right: invertTree(root.Left),
	}
}
