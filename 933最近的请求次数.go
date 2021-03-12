package main

import . "go-leetcode/kit"

type RecentCounter struct {
	q *Queue
}

func Constructor() RecentCounter {
	return RecentCounter{q: NewQueue()}
}

func (this *RecentCounter) Ping(t int) int {
	for this.q.Len() > 0 {
		x := this.q.Front().(int)
		if x < t - 3000 {
			this.q.Pop()
		} else {
			break
		}
	}
	this.q.Push(t)
	return this.q.Len()
}
