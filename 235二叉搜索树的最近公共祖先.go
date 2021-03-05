package main

import (
	"fmt"
	. "go-leetcode/kit"
)

func main() {
	tNode1 := &TreeNode{Val: 2}
	tNode4 := &TreeNode{Val: 10}
	tNode2 := &TreeNode{Val: 8, Right: tNode4}
	tNode3 := &TreeNode{Val: 6, Left: tNode1, Right: tNode2}
	fmt.Println(lowestCommonAncestor(tNode3, tNode2, tNode4))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var arr1 []*TreeNode
	var arr2 []*TreeNode
	searchAncestor(root, p, &arr1)
	searchAncestor(root, q, &arr2)

	// for i := range arr1 {
	// 	fmt.Print(arr1[i].Val, " ")
	// }
	// fmt.Println("")
	// for i := range arr2 {
	// 	fmt.Print(arr2[i].Val, " ")
	// }
	// fmt.Println("")

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return arr1[i]
			}
		}
	}
	return nil
}

func searchAncestor(root, x *TreeNode, arr *[]*TreeNode) bool {
	if root == x {
		*arr = append(*arr, root)
		return true
	}
	if root == nil {
		return false
	}
	if searchAncestor(root.Left, x, arr) || searchAncestor(root.Right, x, arr) {
		*arr = append(*arr, root)
		return true
	}
	return false
}
