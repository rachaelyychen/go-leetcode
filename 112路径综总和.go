package main

import . "go-leetcode/kit"

func hasPathSum(root *TreeNode, targetSum int) bool {
	return InOrder112(root, targetSum, 0)
}

func InOrder112(root *TreeNode, targetSum, sum int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if IsLeaf(root) {
		return targetSum == sum
	}
	return InOrder112(root.Left, targetSum, sum) || InOrder112(root.Right, targetSum, sum)
}
