package kit

type Stack struct {
	arr []interface{}
}

// NewStack 返回 *kit.Stack
func NewStack() *Stack {
	return &Stack{arr: []interface{}{}}
}

// Push 把 n 放入 栈
func (s *Stack) Push(n interface{}) {
	s.arr = append(s.arr, n)
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack) Pop() interface{} {
	if len(s.arr) == 0 {
		return nil
	}
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return res
}

// Len 返回 s 的长度
func (s *Stack) Len() int {
	return len(s.arr)
}

// IsEmpty 返回 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() interface{} {
	if len(s.arr) == 0 {
		return nil
	}
	return s.arr[len(s.arr)-1]
}
