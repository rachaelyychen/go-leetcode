package main

import . "go-leetcode/kit"

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preOrder144(root, &res)
	return res
}

func preOrder144(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	preOrder144(root.Left, arr)
	preOrder144(root.Right, arr)
}
