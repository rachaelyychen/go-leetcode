package main

import . "go-leetcode/kit"

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, arr := make([][]int, 0), make([]int, 0)
	getPathSum(root, target, 0, arr, &res)
	return res
}

func getPathSum(t *TreeNode, target, sum int, arr []int, res *[][]int) {
	arr = append(arr, t.Val)
	if IsLeaf(t) && target == sum+t.Val {
		t := make([]int, len(arr))
		copy(t, arr)
		*res = append(*res, t)
		return
	}
	if t.Left != nil {
		getPathSum(t.Left, target, sum+t.Val, arr, res)
	}
	if t.Right != nil {
		getPathSum(t.Right, target, sum+t.Val, arr, res)
	}
}
