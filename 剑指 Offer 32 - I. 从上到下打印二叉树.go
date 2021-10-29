package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/27/21 9:35 AM
**/

func levelOrder(root *TreeNode) []int {
	var res []int
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			if tNode != nil {
				res = append(res, tNode.Val)
				if tNode.Left != nil {
					q.Push(tNode.Left)
				}
				if tNode.Right != nil {
					q.Push(tNode.Right)
				}
			}
		}
	}
	return res
}
