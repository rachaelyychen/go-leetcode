package temp_pkg

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
* @create: 10/18/21 4:01 PM
**/

// 输入一个链表的头节点，从尾到头打印该链表的节点值。
// 时间复杂度O(n)，空间复杂度O(n)的解法：顺序遍历链表，节点值依次存入一个数组，从后到前打印该数组的元素即可。

func main() {
	node3 := ListNode{Val: 3}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	reversePrintLinkedList(&node1)
}

func reversePrintLinkedList(head *ListNode) {
	var arr []int
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
