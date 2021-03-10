package main

import (
	. "go-leetcode/kit"
)

func middleNode(head *ListNode) *ListNode {
	p, q := head, head.Next
	for q != nil {
		p = p.Next
		if q.Next == nil {
			break
		}
		q = q.Next.Next
	}
	return p
}


