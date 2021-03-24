package main

import . "go-leetcode/kit"

func getDecimalValue(head *ListNode) int {
	var arr []int
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	var res int
	t := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i] * t
		t *= 2
	}
	return res
}
