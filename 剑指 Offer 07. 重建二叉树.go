package main

import (
	. "go-leetcode/kit"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/18/21 5:20 PM
**/

// 根据二叉树前序和中序遍历的结果，重建该二叉树，假设遍历结果都不包含重复数字。
// 前序+中序 / 后序+中序，都可以唯一确定一棵二叉树。固定解法是递归重建，从前序/后序结果中得到根节点，从中序结果中得到左右子树。

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	var rootIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) > 1 {
		root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
		if rootIndex < len(inorder)-1 {
			root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
		}
	}
	return root
}
