package main

import . "go-leetcode/kit"

func isBalanced(root *TreeNode) bool {
	return height(root) > -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	balance := leftHeight - rightHeight
	if leftHeight > -1 && rightHeight > -1 && balance >= -1 && balance <= 1 {
		if balance < 0 {
			return rightHeight + 1
		} else {
			return leftHeight + 1
		}
	}
	return -1
}
