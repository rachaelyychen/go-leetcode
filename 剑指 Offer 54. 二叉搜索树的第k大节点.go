package main

import . "go-leetcode/kit"

func kthLargest(root *TreeNode, k int) int {
	var arr []int
	var inOrder func(t *TreeNode)
	inOrder = func(t *TreeNode) {
		if t == nil {
			return
		}
		if t.Left != nil {
			inOrder(t.Left)
		}
		arr = append(arr, t.Val)
		if t.Right != nil {
			inOrder(t.Right)
		}
	}
	inOrder(root)
	if len(arr) < k {
		return 0
	}
	return arr[len(arr)-k]
}
