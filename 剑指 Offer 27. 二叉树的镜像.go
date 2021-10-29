package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 9:57 AM
**/

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	t := root.Right
	root.Right = mirrorTree(root.Left)
	root.Left = mirrorTree(t)
	return root
}
