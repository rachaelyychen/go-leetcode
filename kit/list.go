package kit

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (self *ListNode) Print() {
	header := self
	for ; header != nil; {
		fmt.Println(header.Val)
		header = header.Next
	}
}
