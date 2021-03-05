package main

import . "go-leetcode/kit"

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	p := newHead
	for head != nil {
		if head.Val != val {
			p.Next = head
			p = p.Next
		}
		head = head.Next
	}
	p.Next = nil
	return newHead.Next
}
