package main

import . "go-leetcode/kit"

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postOrder145(root, &res)
	return res
}

func postOrder145(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	postOrder145(root.Left, arr)
	postOrder145(root.Right, arr)
	*arr = append(*arr, root.Val)
}
