package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	lNode1 := &ListNode{}
	lNode1.Next = lNode1
	fmt.Println(hasCycle(lNode1))
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
				if fast == nil || slow == nil {
					break
				}
			}
			return true
		}
	}
	return false
}
