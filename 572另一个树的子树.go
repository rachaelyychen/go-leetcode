package main

import . "go-leetcode/kit"

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t != nil || s != nil && t == nil {
		return false
	}
	if s == nil && t == nil {
		return true
	}
	// 同一个树或左右子树
	return isSametree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSametree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s != nil && t != nil {
		return s.Val == t.Val && isSametree(s.Left, t.Left) && isSametree(s.Right, t.Right)
	}
	return false
}
