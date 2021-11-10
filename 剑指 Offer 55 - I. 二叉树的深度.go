package main

import . "go-leetcode/kit"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var level int
	var levelOrder func(t *TreeNode)
	levelOrder = func(t *TreeNode) {
		q := NewQueue()
		q.Push(root)
		for q.Len() > 0 {
			l := q.Len()
			for i := 0; i < l; i++ {
				tNode := q.Pop().(*TreeNode)
				if tNode.Left != nil {
					q.Push(tNode.Left)
				}
				if tNode.Right != nil {
					q.Push(tNode.Right)
				}
			}
			level += 1
		}
	}
	levelOrder(root)
	return level
}
