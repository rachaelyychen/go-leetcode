package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 12:38 PM
**/

func deleteNode(head *ListNode, val int) *ListNode {
	p := head
	if p == nil {
		return nil
	}
	if p.Val == val {
		return p.Next
	}
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}
