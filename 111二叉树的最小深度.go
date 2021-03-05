package main

import (
	. "go-leetcode/kit"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		res += 1
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			if IsLeaf(tNode) {
				return res
			}
			if tNode.Left != nil {
				q.Push(tNode.Left)
			}
			if tNode.Right != nil {
				q.Push(tNode.Right)
			}
		}
	}
	return res
}
