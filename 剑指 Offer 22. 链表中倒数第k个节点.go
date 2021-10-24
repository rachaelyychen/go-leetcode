package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/23/21 4:23 PM
**/

// 链表只需遍历一次的解法：先把第一个指针移动到第k个节点，这里可以判断链表是否有k个节点，再移动到第k+1个节点，
// 然后第二个指针放在第1个节点，两指针同时移动直到第一个指针为空，第二个指针就在倒数第k个节点。

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, q := head, head
	cnt := 1
	for p != nil {
		if cnt == k {
			break
		}
		cnt += 1
		p = p.Next
	}
	if p == nil {
		return nil
	}
	p = p.Next
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}
