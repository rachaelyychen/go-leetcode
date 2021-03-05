package main

import (
	"fmt"
	"go-leetcode/kit"
)

func main() {
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
	stack := kit.NewStack()
	arr := []rune(s)
	for i := range arr {
		s := string(arr[i])
		switch s {
		case "(":
			stack.Push(s)
		case "{":
			stack.Push(s)
		case "[":
			stack.Push(s)
		case ")":
			if stack.IsEmpty() || "(" != stack.Pop().(string) {
				return false
			}
		case "}":
			if stack.IsEmpty() || "{" != stack.Pop().(string) {
				return false
			}
		case "]":
			if stack.IsEmpty() || "[" != stack.Pop().(string) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
