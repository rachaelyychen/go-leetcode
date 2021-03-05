package kit

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IntMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	if len(*h) == 0 {
		return nil
	}
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
