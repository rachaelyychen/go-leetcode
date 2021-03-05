package main

func preorder(root *Node) []int {
	var res []int
	preOrder589(root, &res)
	return res
}

func preOrder589(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for i := range root.Children {
		preOrder589(root.Children[i], res)
	}
	return
}
