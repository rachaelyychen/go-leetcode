package main

type MyHashMap struct {
	kvs []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	kvs := make([]int, 1e6+1)
	for i := range kvs {
		kvs[i] = -1
	}
	return MyHashMap{kvs:kvs}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	this.kvs[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.kvs[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	this.kvs[key] = -1
}
