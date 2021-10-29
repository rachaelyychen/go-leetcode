package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 11:39 AM
**/

type MinStack struct {
	arr []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		curMin := this.min[len(this.min)-1]
		if x < curMin {
			curMin = x
		}
		this.min = append(this.min, curMin)
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) > 0 {
		this.arr = this.arr[:len(this.arr)-1]
	}
	if len(this.min) > 0 {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.arr) > 0 {
		return this.arr[len(this.arr)-1]
	}
	return 0
}

func (this *MinStack) Min() int {
	if len(this.min) > 0 {
		return this.min[len(this.min)-1]
	}
	return 0
}
