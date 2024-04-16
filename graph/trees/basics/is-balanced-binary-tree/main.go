package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// Define a helper function to perform depth calculation and balance check simultaneously
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		// Base case: If the node is nil, return depth 0 and true for balance
		if node == nil {
			return 0, true
		}

		// Recursively calculate the depths of the left and right subtrees
		leftDepth, isBalancedLeft := dfs(node.Left)
		rightDepth, isBalancedRight := dfs(node.Right)

		// Check if either subtree is unbalanced or if the difference in depths exceeds 1
		if !isBalancedLeft || !isBalancedRight || abs(leftDepth-rightDepth) > 1 {
			return -1, false
		}

		// Calculate the depth of the current subtree
		depth := 1 + max(leftDepth, rightDepth)

		// Return the depth and balance status of the current subtree
		return depth, true
	}

	// Invoke the helper function on the root node and return the balance status
	_, isBalanced := dfs(root)
	return isBalanced
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
