package temp_pkg

import (
	"fmt"
	. "go-leetcode/kit"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/18/21 5:20 PM
**/

// 根据二叉树前序和中序遍历的结果，重建该二叉树，假设遍历结果都不包含重复数字。
// 前序+中序 / 后序+中序，都可以唯一确定一棵二叉树。固定解法是递归重建，从前序/后序结果中得到根节点，从中序结果中得到左右子树。

func main() {
	preOrderRes, inOrderRes := []int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}
	// preOrderRes, inOrderRes := []int{1, 2, 3, 4}, []int{4, 3, 2, 1}
	var res []int
	preOrder(rebuildBinaryTree(preOrderRes, inOrderRes), &res)
	for i := range res {
		if preOrderRes[i] != res[i] {
			fmt.Println("rebuild error")
			return
		}
	}
}

func rebuildBinaryTree(preOrderRes, inOrderRes []int) *TreeNode {
	if len(preOrderRes) == 0 || len(inOrderRes) == 0 || len(preOrderRes) != len(inOrderRes) {
		return nil
	}
	var rootIndex int
	for i := 0; i < len(inOrderRes); i++ {
		if inOrderRes[i] == preOrderRes[0] {
			rootIndex = i
			break
		}
	}
	root := &TreeNode{Val: preOrderRes[0]}
	if len(preOrderRes) > 1 {
		root.Left = rebuildBinaryTree(preOrderRes[1:rootIndex+1], inOrderRes[:rootIndex])
		if rootIndex < len(inOrderRes)-1 {
			root.Right = rebuildBinaryTree(preOrderRes[rootIndex+1:], inOrderRes[rootIndex+1:])
		}
	}
	return root
}

func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
