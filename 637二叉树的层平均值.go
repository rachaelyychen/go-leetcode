package main

import (
	. "go-leetcode/kit"
)

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	levelOrder637(root, &res)
	return res
}

func levelOrder637(root *TreeNode, res *[]float64) {
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		var sum int
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			if tNode != nil {
				sum += tNode.Val
				if tNode.Left != nil {
					q.Push(tNode.Left)
				}
				if tNode.Right != nil {
					q.Push(tNode.Right)
				}
			}
		}
		*res = append(*res, float64(sum)/float64(length))
	}
}
