package main

import . "go-leetcode/kit"

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inOrder94(root, &res)
	return res
}

func inOrder94(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder94(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder94(root.Right, arr)
}
