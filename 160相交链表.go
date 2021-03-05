package main

import (
	. "go-leetcode/kit"
)

func main() {
	lNode1 := &ListNode{
		Val: 1,
	}
	lNode2 := &ListNode {
		Val: 2,
		Next: lNode1,
	}
	lNode3 := &ListNode{
		Val: 3,
		Next: lNode1,
	}
	getIntersectionNode(lNode2, lNode3)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getListLen(headA), getListLen(headB)
	if lenA > lenB {
		for i := 0; i < (lenA - lenB) && headA != nil; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < (lenB - lenA) && headB != nil; i++ {
			headB = headB.Next
		}
	}
	for headA != headB {
		if headA == nil || headB == nil {
			headA = nil
			break
		}
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getListLen(head *ListNode) (length int) {
	p := head
	for p != nil {
		length += 1
		p = p.Next
	}
	return length
}
