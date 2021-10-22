package main

import (
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

func reversePrint(head *ListNode) []int {
	var arr []int
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	var res []int
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}
