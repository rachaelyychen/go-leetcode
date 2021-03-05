package main

import . "go-leetcode/kit"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth + 1
	}
	return rightMaxDepth + 1
}
