package main

import . "go-leetcode/kit"

func main() {
	lNode1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	lNode2 := &ListNode{
		Val:  1,
		Next: lNode1,
	}
	lNode3 := &ListNode{
		Val:  1,
		Next: lNode2,
	}
	deleteDuplicates(lNode3).Print()
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		p := head
		q := p.Next
		for p != nil && q != nil {
			if p.Val == q.Val {
				q = q.Next
			} else {
				p.Next = q
				p = q
				q = q.Next
			}
		}
		p.Next = nil
	}
	return head
}
