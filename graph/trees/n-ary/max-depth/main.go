package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	// Key Idea: Use recursion
	if root == nil {
		return 0
	}

	// At the root level `depth`
	depth := 1
	for _, child := range root.Children {
		// Plus 1 is for the child node that will be treated as the `root` in the next recursive stack frame
		depth = max(depth, 1+maxDepth(child))
	}
	return depth
}
