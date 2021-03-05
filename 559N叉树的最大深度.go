package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var rootDepth int
	for i := range root.Children {
		childDepth := maxDepth(root.Children[i])
		if childDepth > rootDepth {
			rootDepth = childDepth
		}
	}
	return rootDepth+1
}
