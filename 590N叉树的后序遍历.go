package main

func postorder(root *Node) []int {
	var res []int
	postOrder590(root, &res)
	return res
}

func postOrder590(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for i := range root.Children{
		postOrder590(root.Children[i], res)
	}
	*res = append(*res, root.Val)
	return
}
