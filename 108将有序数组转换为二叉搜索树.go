package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	root := sortedArrayToBST([]int{-10,-3,0,5,9})
	if root != nil {
		fmt.Println(root.LevelOrder())
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	m := left + (right - left) / 2
	root := &TreeNode{Val: nums[m]}
	root.Left = buildBST(nums, left, m-1)
	root.Right = buildBST(nums, m+1, right)
	return root
}
