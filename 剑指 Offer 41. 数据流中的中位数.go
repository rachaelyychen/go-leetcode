package main

import (
	"container/heap"
	. "go-leetcode/kit"
)

// 将一段排序后的数据流分为前一半和后一半，前一半用最大堆存放，后一半用最小堆存放，
// 总数为偶数时中位数就是最大堆堆顶和最小堆堆顶的平均值；总数为奇数时中位数是最大堆堆顶。

type MedianFinder struct {
	minHeap *IntMinHeap
	maxHeap *IntMaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{minHeap: NewIntMinHeap(), maxHeap: NewIntMaxHeap()}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 || this.minHeap.Top().(int) > num {
		heap.Push(this.maxHeap, num)
		return
	}
	heap.Push(this.minHeap, num)
}

func (this *MedianFinder) FindMedian() float64 {
	tot := this.maxHeap.Len() + this.minHeap.Len()
	for this.maxHeap.Len() > tot/2+tot&0x1 {
		t := this.maxHeap.Top()
		heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, t)
	}
	for this.minHeap.Len() > tot/2 {
		t := this.minHeap.Top()
		heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, t)
	}
	if tot&0x1 == 1 {
		return (float64)(this.maxHeap.Top().(int))
	}
	return ((float64)(this.maxHeap.Top().(int)) + (float64)(this.minHeap.Top().(int))) / 2
}
