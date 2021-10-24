package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 8:40 AM
**/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	p := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			break
		}
		if l2 == nil {
			p.Next = l1
			break
		}
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	return head.Next
}
