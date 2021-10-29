package main

// Definition for a RandomNode.
type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// 分三步：1、克隆新节点加在原节点后面；2、处理新节点的Random；3、分离原链表和新链表

func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return head
	}
	var p, q *RandomNode
	p = head
	for p != nil {
		q = new(RandomNode)
		q.Val = p.Val
		q.Next = p.Next
		p.Next = q
		p = q.Next
	}
	p, q = head, head.Next
	for p != nil && q != nil {
		if p.Random != nil {
			q.Random = p.Random.Next
		}
		p = q.Next
		if p != nil {
			q = p.Next
		}
	}
	newHead := head.Next
	p, q = head, head.Next
	for p != nil && q != nil {
		p.Next = p.Next.Next
		if q.Next != nil {
			q.Next = q.Next.Next
		}
		p = p.Next
		q = q.Next
	}
	return newHead
}
