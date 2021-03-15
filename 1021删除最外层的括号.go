package main

import . "go-leetcode/kit"

func removeOuterParentheses(S string) string {
	s := NewStack()
	var res string
	for _, b := range []byte(S) {
		if b == ')' {
			s.Pop()
		}
		if !s.IsEmpty() {
			res += string(b)
		}
		if b == '(' {
			s.Push(b)
		}
	}
	return res
}


