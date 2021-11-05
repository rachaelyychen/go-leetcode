package kit

import "container/heap"

// An IntMinHeap is a min-heap of ints.
type IntMinHeap []int

func NewIntMinHeap() *IntMinHeap {
	h := &IntMinHeap{}
	heap.Init(h)
	return h
}

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntMinHeap) Top() interface{} {
	return (*h)[0]
}

// An IntMaxHeap is a max-heap of ints.
type IntMaxHeap []int

func NewIntMaxHeap() *IntMaxHeap {
	h := &IntMaxHeap{}
	heap.Init(h)
	return h
}

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntMaxHeap) Top() interface{} {
	return (*h)[0]
}

// h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
