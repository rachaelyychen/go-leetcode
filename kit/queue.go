package kit

type Queue struct {
	arr []interface{}
}

// NewQueue 返回 *kit.Queue
func NewQueue() *Queue {
	return &Queue{arr: []interface{}{}}
}

// Push 把 n 放入队列
func (q *Queue) Push(n interface{}) {
	q.arr = append(q.arr, n)
}

// Pop 从 q 中取出最先进入队列的值
func (q *Queue) Pop() interface{} {
	if len(q.arr) == 0 {
		return nil
	}
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}

// Len 返回 q 的长度
func (q *Queue) Len() int {
	return len(q.arr)
}

// IsEmpty 反馈 q 是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
