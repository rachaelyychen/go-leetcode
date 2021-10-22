package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/16/21 6:47 PM
**/

// 用两个栈实现一个队列的进队、出队操作。
// 进队操作就是第一个栈s1的进栈；
// 出队操作需要用到第二个栈s2，如果s2为空，就把s1所有元素压进s2，s2栈顶元素就是队头元素，出栈即可；如果s2不为空，直接把s2栈顶元素出栈。

type CQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() CQueue {
	return CQueue{s1: NewStack(), s2: NewStack()}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.IsEmpty() {
		for {
			if this.s1.IsEmpty() {
				break
			}
			this.s2.Push(this.s1.Pop())
		}
	}
	t := -1
	if !this.s2.IsEmpty() {
		t = this.s2.Top().(int)
		this.s2.Pop()
	}
	return t
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
