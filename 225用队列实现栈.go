package main

import "fmt"

func main() {
	obj := MyStackConstructor()
	obj.Push(1)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyStack struct {
	arr []int
}

/** Initialize your data structure here. */
func MyStackConstructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.arr = append(this.arr, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.arr[len(this.arr)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}
