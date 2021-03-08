package main

import . "go-leetcode/kit"

func minDiffInBST(root *TreeNode) int {
	res, preVal := -1, -1
	inOrder783(root, &preVal, &res)
	return res
}

func inOrder783(root *TreeNode, preVal, minDiff *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder530(root.Left, preVal, minDiff)
	}
	if *preVal > 1 {
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
		inOrder783(root.Right, preVal, minDiff)
	}
}

