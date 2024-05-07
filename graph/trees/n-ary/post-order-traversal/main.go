package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {

	// Key Idea: Process child nodes first recursively and then append the root.
	// Opposite of pre-order traversal.

	var result []int
	if root == nil {
		return result
	}

	if len(root.Children) == 0 {
		return []int{root.Val}
	}

	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}

	result = append(result, root.Val)
	return result
}
