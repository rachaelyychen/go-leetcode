package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(1)
	fmt.Println(obj.Contains(1))
	obj.Remove(1)
	fmt.Println(obj.Contains(1))
}

type MyHashSet struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]int, 1e6+1)}
}

func (this *MyHashSet) Add(key int)  {
	if this.arr[key] == 0 {
		this.arr[key] = 1
	}
}

func (this *MyHashSet) Remove(key int)  {
	if this.arr[key] == 1 {
		this.arr[key] = 0
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.arr[key] == 1
}
