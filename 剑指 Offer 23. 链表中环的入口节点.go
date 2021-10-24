package main

import (
	"fmt"
	. "go-leetcode/kit"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/23/21 4:48 PM
**/

// 找到链表中环的入口节点分为以下两步：
// 1、确定链表中是否存在环，如果存在，环的节点数是多少。
// 使用一快一慢两个指针从头节点开始同时遍历，快指针每次移动两步，慢指针每次移动一步，第一次相遇在非空节点说明链表中存在环，而且相遇节点必定在环中，使用慢指针继续遍历同时计数，回到相遇节点后即可得到环的节点数。
// 2、找到环的入口节点。
// 使用两个指针从头节点开始遍历，第一个指针先移动步数等于环的节点数，然后两指针同时移动，第一次相遇的节点就是环的入口节点。

func entryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil || head.Next == head {
		return head
	}
	p, q := head, head
	for p != nil && q != nil {
		p = p.Next
		if q.Next == nil {
			q = q.Next
		} else {
			q = q.Next.Next
		}
		if p != nil && p == q {
			break
		}
	}
	if p == nil || q == nil {
		return nil
	}
	var loopCnt int // 存在环，loopCnt是大于0的
	for {
		p = p.Next
		loopCnt += 1
		if p == q {
			break
		}
	}
	p, q = head, head
	for p != nil {
		p = p.Next
		loopCnt -= 1
		if loopCnt == 0 {
			break
		}
	}
	for p != nil && q != nil {
		if p == q { // 判断条件，头节点可能是环入口
			break
		}
		p = p.Next
		q = q.Next
	}
	return p
}

func main() {
	node1, node2, node3 := ListNode{Val: 1}, ListNode{Val: 2}, ListNode{Val: 3}
	node2.Next = &node3
	node1.Next = &node2
	node3.Next = &node1
	fmt.Println(entryNodeOfLoop(&node1))
}
