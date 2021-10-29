package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/27/21 8:58 PM
**/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		var t []int
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			if tNode != nil {
				t = append(t, tNode.Val)
				if tNode.Left != nil {
					q.Push(tNode.Left)
				}
				if tNode.Right != nil {
					q.Push(tNode.Right)
				}
			}
		}
		res = append(res, t)
	}
	return res
}
