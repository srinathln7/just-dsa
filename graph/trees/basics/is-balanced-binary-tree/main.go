package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var isBalanced func(node *TreeNode) bool
	isBalanced = func(node *TreeNode) bool {
		// Base case: Empty trees are always balanced
		if node == nil {
			return true
		}

		leftDepth, rightDepth := depth(node.Left), depth(node.Right)
		return isBalanced(node.Left) && abs(leftDepth-rightDepth) <= 1 && isBalanced(node.Right)
	}

	return isBalanced(root)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	leftDepth, rightDepth := depth(node.Left), depth(node.Right)

	// Plus one is for the root node
	return 1 + max(leftDepth, rightDepth)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
