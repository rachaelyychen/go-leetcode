package main

import (
	"fmt"
	. "go-leetcode/kit"
)

// BST的中序遍历结果是排序数组，转成链表的每个节点左指针指向前驱（左子树的最大节点），右指针指向后继（右子树的最小节点）。
// 使用treeCurrentNode指向BST中序遍历的当前节点，listLastNode指向双向链表的最后一个节点，那么listLastNode就是已转换好的左子树最大节点，
// 把treeCurrentNode加入链表中，继续遍历过程。
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, tail *TreeNode
	convertNode(root, &head)
	tail = head
	for head != nil && head.Left != nil {
		head = head.Left
	}
	tail.Right = head
	head.Left = tail
	return head
}

func convertNode(treeCurrentNode *TreeNode, listLastNode **TreeNode) {
	if treeCurrentNode == nil {
		return
	}
	convertNode(treeCurrentNode.Left, listLastNode)
	treeCurrentNode.Left = *listLastNode
	if *listLastNode != nil {
		(*listLastNode).Right = treeCurrentNode
	}
	*listLastNode = treeCurrentNode
	convertNode(treeCurrentNode.Right, listLastNode)
}

func main() {
	t1, t2 := TreeNode{Val: 1}, TreeNode{Val: 3}
	root := TreeNode{Val: 2, Left: &t1, Right: &t2}
	head := treeToDoublyList(&root)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Right
	}
}
