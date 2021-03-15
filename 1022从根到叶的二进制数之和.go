package main

import . "go-leetcode/kit"

func sumRootToLeaf(root *TreeNode) int {
	var res int
	preOrder1022(root, 0, &res)
	return res
}

func preOrder1022(root *TreeNode, num int, res *int) {
	if root == nil {
		return
	}
	num = num<<1 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += num
	}
	if root.Left != nil {
		preOrder1022(root.Left, num, res)
	}
	if root.Right != nil {
		preOrder1022(root.Right, num, res)
	}
}
