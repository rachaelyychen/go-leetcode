package main

import . "go-leetcode/kit"

func sumOfLeftLeaves(root *TreeNode) int {
	var res int
	inOrder404(root, &res, false)
	return res
}

func inOrder404(root *TreeNode, res *int, isLeft bool) {
	if root == nil {
		return
	}
	if isLeft && root.Left == nil && root.Right == nil {
		*res += root.Val
	}
	if root.Left != nil {
		inOrder404(root.Left, res, true)
	}
	if root.Right != nil {
		inOrder404(root.Right, res, false)
	}
}
