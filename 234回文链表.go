package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	lNode1 := &ListNode{Val:1}
	lNode2 := &ListNode{Val:2, Next:lNode1}
	lNode3 := &ListNode{Val:1, Next:lNode2}
	fmt.Println(isPalindrome234(lNode3))
}

func isPalindrome234(head *ListNode) bool {
	length, p := 0, head
	for p != nil {
		length += 1
		p = p.Next
	}
	if length == 0 {
		return true
	}
	var arr []int
	p = head
	cnt := 0
	for cnt < length/2 {
		arr = append(arr, p.Val)
		cnt += 1
		p = p.Next
	}
	if length%2 > 0 {
		p = p.Next
	}
	i := len(arr) - 1
	for p != nil && i > -1 {
		if p.Val != arr[i] {
			return false
		}
		p = p.Next
		i -= 1
	}
	return true
}
