package main

import . "go-leetcode/kit"

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	postOrder543(root, &res)
	return res
}

func postOrder543(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := postOrder543(root.Left, res), postOrder543(root.Right, res)
	diameter := leftHeight+rightHeight
	if diameter > *res {
		*res = diameter
	}
	if leftHeight > rightHeight {
		return leftHeight+1
	} else {
		return rightHeight+1
	}
}
