package main

import . "go-leetcode/kit"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, p, q *ListNode
	l3 = new(ListNode)
	p = l3
	var t int
	for l1 != nil || l2 != nil {
		q = new(ListNode)
		var s int
		if l1 == nil && l2 != nil {
			s = l2.Val + t
			l2 = l2.Next
		}
		if l2 == nil && l1 != nil {
			s = l1.Val + t
			l1 = l1.Next
		}
		if l1 != nil && l2 != nil {
			s = l1.Val + l2.Val + t
			l1 = l1.Next
			l2 = l2.Next
		}
		t = s / 10
		q.Val = s % 10
		p.Next = q
		p = q
	}
	if t != 0 {
		p.Next = &ListNode{Val: t}
	}
	return l3.Next
}
