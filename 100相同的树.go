package main

import (
	"fmt"
	. "go-leetcode/kit"
	"strconv"
)

func main() {
	// 	1 null 1
	t1Node1 := &TreeNode{
		Val: 1,
	}
	t1Node2 := &TreeNode{
		Val:   1,
		Right: t1Node1,
	}
	// 1 1
	t2Node1 := &TreeNode{
		Val: 1,
	}
	t2Node2 := &TreeNode{
		Val:   1,
		Left:  t2Node1,
	}
	fmt.Println(isSameTree(t1Node2, t2Node2))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var rp []string
	InOrder100(p, &rp)
	var rq []string
	InOrder100(q, &rq)

	// fmt.Println(rp, rq)

	if len(rp) != len(rq) {
		return false
	}
	for i := 0; i < len(rp); i++ {
		if rp[i] != rq[i] {
			return false
		}
	}
	return true
}

func InOrder100(t *TreeNode, res *[]string) {
	if t == nil {
		*res = append(*res, "null")
		return
	}
	*res = append(*res, strconv.Itoa(t.Val))
	if t.Left != nil || t.Right != nil {
		InOrder100(t.Left, res)
	}
	if t.Right != nil {
		InOrder100(t.Right, res)
	}
}
