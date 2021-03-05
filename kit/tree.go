package kit

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) LevelOrder() []string {
	var res []string
	q := NewQueue()
	q.Push(t)
	for !q.IsEmpty() {
		length := q.Len()
		for i := 0; i < length; i++ {
			tNode := q.Pop().(*TreeNode)
			if tNode == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(tNode.Val))
				if tNode.Left != nil || tNode.Right != nil {
					q.Push(tNode.Left)
					q.Push(tNode.Right)
				}
			}
		}
	}
	return res
}

func IsLeaf(t *TreeNode) bool {
	if t == nil {
		return false
	}
	return t.Left == nil && t.Right == nil
}
