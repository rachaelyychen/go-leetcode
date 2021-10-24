package main

import . "go-leetcode/kit"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 8:51 AM
**/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return A.Val == B.Val && isEqualTree(A.Left, B.Left) && isEqualTree(A.Right, B.Right) ||
		isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isEqualTree(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A != nil && A.Val == B.Val && isEqualTree(A.Left, B.Left) && isEqualTree(A.Right, B.Right) {
		return true
	}
	return false
}
