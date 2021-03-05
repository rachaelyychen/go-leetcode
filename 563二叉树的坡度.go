package main

import . "go-leetcode/kit"

func findTilt(root *TreeNode) int {
	var res int
	findTilt563(root, &res)
	return res
}

func findTilt563(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftSum, rightSum, tilt := findTilt563(root.Left, res), findTilt563(root.Right, res), 0
	if leftSum > rightSum {
		tilt = leftSum - rightSum
	} else {
		tilt = rightSum - leftSum
	}
	*res += tilt
	return leftSum+rightSum+root.Val
}
