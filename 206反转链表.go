package main

import . "go-leetcode/kit"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p *ListNode
	q := head
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}
	return p
}
