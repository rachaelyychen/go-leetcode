package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	// 2 2 null
	tNode1 := &TreeNode{
		Val: 2,
	}
	tNode2 := &TreeNode{
		Val:  2,
		Left: tNode1,
	}
	fmt.Println(isSymmetric(tNode2))
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
	}
	return false
}
