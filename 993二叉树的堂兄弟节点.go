package main

import . "go-leetcode/kit"

func isCousins(root *TreeNode, x int, y int) bool {
	xLevel, yLevel, xFather, yFather := levelOrder993(root, x, y)
	if xLevel == yLevel && xFather != yFather {
		return true
	}
	return false
}

func levelOrder993(root *TreeNode, x, y int) (xLevel, yLevel, xFather, yFather int) {
	if root == nil {
		return
	}
	type node1 struct {
		*TreeNode
		father int
	}
	q := NewQueue()
	q.Push(node1{root, 0})
	level := 0
	for !q.IsEmpty() {
		length := q.Len()
		for i := 0; i < length; i++ {
			t := q.Pop().(node1)
			if t.Val == x {
				xLevel = level
				xFather = t.father
			}
			if t.Val == y {
				yLevel = level
				yFather = t.father
			}
			if t.Left != nil {
				q.Push(node1{t.Left, t.Val})
			}
			if t.Right != nil {
				q.Push(node1{t.Right, t.Val})
			}
		}
		level += 1
	}
	return
}


