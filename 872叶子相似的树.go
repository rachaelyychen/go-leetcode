package main

import . "go-leetcode/kit"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leafs1, leafs2 []int
	preOrder872(root1, &leafs1)
	preOrder872(root2, &leafs2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := range leafs1 {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}

func preOrder872(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}
	if root.Left != nil {
		preOrder872(root.Left, leafs)
	}
	if root.Right != nil {
		preOrder872(root.Right, leafs)
	}
}


