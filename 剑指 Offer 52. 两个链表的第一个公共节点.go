package main

import . "go-leetcode/kit"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	p, q := headA, headB
	for p != nil {
		lenA += 1
		p = p.Next
	}
	for q != nil {
		lenB += 1
		q = q.Next
	}
	p, q = headA, headB
	var cnt int
	if lenA > lenB {
		cnt = lenA - lenB
		for p != nil {
			if cnt == 0 {
				break
			}
			cnt -= 1
			p = p.Next
		}
	} else {
		cnt = lenB - lenA
		for q != nil {
			if cnt == 0 {
				break
			}
			cnt -= 1
			q = q.Next
		}
	}
	for p != q && p != nil && q != nil {
		p = p.Next
		q = q.Next
	}
	if p == nil || q == nil {
		return nil
	}
	return p
}
