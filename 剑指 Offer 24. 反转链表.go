package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 7:02 AM
**/

// 三指针法反转链表。

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
