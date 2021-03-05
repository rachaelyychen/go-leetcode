package main

import (
	. "go-leetcode/kit"
)

func getMinimumDifference(root *TreeNode) int {
	res, preVal := -1, -1
	inOrder530(root, &preVal, &res)
	return res
}

func inOrder530(root *TreeNode, preVal, minDiff *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder530(root.Left, preVal, minDiff)
	}
	if *preVal >= 0 {
		var diff int
		if root.Val > *preVal {
			diff = root.Val - *preVal
		} else {
			diff = *preVal - root.Val
		}
		if *minDiff >= 0 && diff < *minDiff {
			*minDiff = diff
		}
		if *minDiff < 0 {
			*minDiff = diff
		}
	}
	*preVal = root.Val
	if root.Right != nil {
		inOrder530(root.Right, preVal, minDiff)
	}
}
