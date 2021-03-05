package main

import "fmt"

func main() {
	obj := MyQueueConstructor()
	obj.Push(1)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyQueue struct {
	arr []int
}


/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.arr = append(this.arr, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var x int
	if len(this.arr) > 0 {
		x = this.arr[0]
		this.arr = this.arr[1:len(this.arr)]
	}
	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	var x int
	if len(this.arr) > 0 {
		x = this.arr[0]
	}
	return x
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}
