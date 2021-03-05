package main

import . "go-leetcode/kit"

func findTarget(root *TreeNode, k int) bool {
	var res bool
	var arr []int
	inOrder653(root, k, &arr, &res)
	return res
}

func inOrder653(root *TreeNode, k int, arr *[]int, res *bool) {
	if root == nil || *res == true {
		return
	}
	if root.Left != nil {
		inOrder653(root.Left, k, arr, res)
	}
	for i := range *arr {
		if (*arr)[i] + root.Val == k {
			*res = true
			break
		}
	}
	if *res == false {
		*arr = append(*arr, root.Val)
	}
	if root.Right != nil {
		inOrder653(root.Right, k, arr, res)
	}
}
