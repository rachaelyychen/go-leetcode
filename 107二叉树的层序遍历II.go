package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	tNode1 := &TreeNode{
		Val: 9,
	}
	tNode2 := &TreeNode{
		Val: 20,
	}
	tNode3 := &TreeNode{
		Val:   3,
		Left:  tNode1,
		Right: tNode2,
	}
	fmt.Println(levelOrderBottom(tNode3))
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		var level []int
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			level = append(level, tNode.Val)
			if tNode.Left != nil {
				q.Push(tNode.Left)
			}
			if tNode.Right != nil {
				q.Push(tNode.Right)
			}
		}
		res = append(res, level)
	}
	var newRes [][]int
	for i := len(res)-1; i >= 0; i-- {
		newRes = append(newRes, res[i])
	}
	return newRes
}
