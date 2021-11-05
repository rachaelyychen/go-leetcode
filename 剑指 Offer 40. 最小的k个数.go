package main

import (
	"container/heap"
	. "go-leetcode/kit"
)

// 时间复杂度O(nlogk)，空间复杂度O(k)的解法：使用最大堆，遍历数组，当堆中数字不满k个时插入，满k个时比较堆中最大值和当前遍历数字的大小，小于最大值时替换最大值。

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	h := NewIntMaxHeap()
	for i := range arr {
		if h.Len() < k {
			heap.Push(h, arr[i])
		} else {
			if h.Top().(int) > arr[i] {
				heap.Pop(h)
				heap.Push(h, arr[i])
			}
		}
	}
	var res []int
	for h.Len() > 0 {
		t := heap.Pop(h).(int)
		res = append(res, t)
	}
	return res
}
