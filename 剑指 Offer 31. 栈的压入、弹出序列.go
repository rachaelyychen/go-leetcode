package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/26/21 9:16 PM
**/

// 遍历弹出序列的每个元素，如果栈顶元素相等，则弹出栈顶元素，
// 如果栈顶元素不等，则把压入序列的元素依次进栈，直到栈顶元素相等，弹出栈顶元素。

func validateStackSequences(pushed []int, popped []int) bool {
	if len(popped) == 0 {
		return true
	}
	var index int
	s := NewStack()
	for i := range popped {
		if s.IsEmpty() || s.Top().(int) != popped[i] {
			for index < len(pushed) {
				if pushed[index] == popped[i] {
					if index < len(pushed)-1 {
						index += 1
					}
					break
				}
				s.Push(pushed[index])
				index += 1
			}
			if index == len(pushed) {
				return false
			}
			continue
		}
		s.Pop()
	}
	if !s.IsEmpty() {
		return false
	}
	return true
}
