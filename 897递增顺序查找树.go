package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	tNode1 := &TreeNode{Val: 0}
	tNode2 := &TreeNode{Val: 1}
	tNode3 := &TreeNode{Val: 2, Left: tNode1, Right: tNode2}
	fmt.Println(increasingBST(tNode3))
}

var res, cur *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	inOrder897(root)
	return res
}

func inOrder897(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder897(root.Left)
	}
	if res == nil {
		res = &TreeNode{Val: root.Val}
		cur = res
	} else if cur != nil {
		cur.Right = &TreeNode{Val: root.Val}
		cur = cur.Right
	}
	if root.Right != nil {
		inOrder897(root.Right)
	}
}
