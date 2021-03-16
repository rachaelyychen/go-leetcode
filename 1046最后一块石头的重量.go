package main

import (
	"container/heap"
	. "go-leetcode/kit"
)

func lastStoneWeight(stones []int) int {
	h := NewIntMaxHeap()
	for _, stone := range stones {
		heap.Push(h, stone)
	}
	for h.Len() > 1 {
		s1, s2 := heap.Pop(h).(int), heap.Pop(h).(int)
		if s1-s2 != 0 {
			heap.Push(h, s1-s2)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
