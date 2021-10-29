package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/27/21 10:58 PM
**/

// BST的后序遍历结果，最后一个元素必定是根节点，根据根节点划分左右子树（左子树都小于等于根节点，右子树都大于等于根节点），再做递归判断。

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	rootVal, pivot := postorder[len(postorder)-1], -1
	for i := 0; i < len(postorder)-1; i++ {
		if i+1 < len(postorder)-1 {
			if postorder[i] <= rootVal && postorder[i+1] <= rootVal || postorder[i] >= rootVal && postorder[i+1] >= rootVal {
				continue
			}
			if postorder[i] <= rootVal && postorder[i+1] >= rootVal {
				pivot = i + 1
				continue
			}
			return false
		}
	}
	if pivot == -1 {
		return verifyPostorder(postorder[:len(postorder)-1])
	}
	return verifyPostorder(postorder[:pivot]) && verifyPostorder(postorder[pivot:])
}
