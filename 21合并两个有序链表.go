package main

import . "go-leetcode/kit"

func main() {
	l1Node1 := ListNode{
		Val:  4,
		Next: nil,
	}
	l1Node2 := ListNode{
		Val:  2,
		Next: &l1Node1,
	}
	l2Node1 := ListNode{
		Val:  3,
		Next: nil,
	}
	l2Node2 := ListNode{
		Val:  1,
		Next: &l2Node1,
	}
	// l1: 2->4, l2: 1->3
	mergeTwoLists(&l1Node2, &l2Node2).Print()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	header := res
	for ; l1 != nil || l2 != nil; {
		if l1 == nil {
			header.Next = l2
			break
		}
		if l2 == nil {
			header.Next = l1
			break
		}
		if l1.Val < l2.Val {
			header.Next = l1
			l1 = l1.Next
		} else {
			header.Next = l2
			l2 = l2.Next
		}
		header = header.Next
	}
	return res.Next
}
