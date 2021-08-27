package main

import . "go-leetcode/kit"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt, t := 0, 0
	p := head
	for p != nil {
		cnt += 1
		p = p.Next
	}
	if cnt < n {
		return nil
	}
	if cnt == n {
		return head.Next
	}
	p = head
	for t < cnt-n-1 && p != nil {
		t += 1
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	}
	return head
}
