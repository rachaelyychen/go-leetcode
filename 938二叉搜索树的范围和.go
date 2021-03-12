package main

import . "go-leetcode/kit"

func rangeSumBST(root *TreeNode, low int, high int) int {
	var res int
	preOrder938(root, low, high, &res)
	return res
}

func preOrder938(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}
	if root.Val >= low && root.Val <= high {
		*res += root.Val
	}
	if root.Left != nil {
		preOrder938(root.Left, low, high, res)
	}
	if root.Right != nil {
		preOrder938(root.Right, low, high, res)
	}
}


