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
* @create: 10/27/21 10:45 PM
**/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := NewQueue()
	q.Push(root)
	var level int
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
		level += 1
		if level&0x1 == 0 {
			i, j := 0, len(t)-1
			for i < j {
				x := t[i]
				t[i] = t[j]
				t[j] = x
				i += 1
				j -= 1
			}
		}
		res = append(res, t)
	}
	return res
}
