package main

import . "go-leetcode/kit"

// 时间复杂度O(h+k),中序遍历非递归

func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	s := NewStack()
	p := root
	for p != nil || !s.IsEmpty() {
		for p != nil {
			s.Push(p)
			p = p.Left
		}
		t := s.Top().(*TreeNode)
		arr = append(arr, t.Val)
		if len(arr) == k {
			return arr[k-1]
		}
		s.Pop()
		p = t.Right
	}
	return -1
}

// 时间复杂度O(n),中序遍历递归

func kthSmallest2(root *TreeNode, k int) int {
	var arr []int
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		arr = append(arr, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return arr[k-1]
}
