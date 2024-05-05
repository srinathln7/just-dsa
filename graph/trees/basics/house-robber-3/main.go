package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	// Key Idea: At any point, we have two decisions to make - either include or exclude the current node in our robbery process
	// Map the decision tree

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {

		// Base case:
		if node == nil {
			return 0, 0
		}

		// Recursively calculate for both left and right subtrees
		leftIncl, leftExcl := dfs(node.Left)
		rightIncl, rightExcl := dfs(node.Right)

		// Decision to include the current node which means the nodes immediate children cannot be included
		incl := node.Val + leftExcl + rightExcl

		// Decision to exclude the current node which means the nodes immediate children may or may not be included
		excl := max(leftIncl, leftExcl) + max(rightIncl, rightExcl)

		return incl, excl
	}

	incl, excl := dfs(root)
	return max(incl, excl)
}
