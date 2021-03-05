package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	tNode1 := &TreeNode{
		Val:1,
	}
	tNode2 := &TreeNode {
		Val:2,
		Left: tNode1,
	}
	fmt.Println(findMode(tNode2))
}

func findMode(root *TreeNode) []int {
	// BST的中序遍历是一个升序序列，如果有相等的数，是1 2 2 2 3 3这样的遍历顺序
	var res []int
	isPreNil, preVal, preCnt, maxCnt := true, 0, 0, 0
	inOrder501(root, &isPreNil, &preVal, &preCnt, &maxCnt, &res)
	return res
}

func inOrder501(root *TreeNode, isPreNil *bool, preVal, preCnt, maxCnt *int, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder501(root.Left, isPreNil, preVal, preCnt, maxCnt, res)
	}
	if *isPreNil == false && *preVal == root.Val {
		*preCnt += 1
	} else {
		*preCnt = 1
	}
	if *preCnt == *maxCnt {
		*res = append(*res, root.Val)
	} else if *preCnt > *maxCnt {
		*res = []int{root.Val}
		*maxCnt = *preCnt
	}

	fmt.Print("now node:", root.Val)
	if *isPreNil == false {
		fmt.Print("\tpre node:", *preVal)
	}
	fmt.Println("\tres:", *res)

	*isPreNil = false
	*preVal = root.Val
	if root.Right != nil {
		inOrder501(root.Right, isPreNil, preVal, preCnt, maxCnt, res)
	}
}
