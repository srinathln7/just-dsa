package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	var result []int
	if root == nil {
		return result
	}

	// pre-order => root -> left -> right
	result = []int{root.Val}
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}

	return result
}
