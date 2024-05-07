package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

	// Key Idea: Perform a DFS and calculate its sum when we reach the leaf node
	// As we progress along the path, every time the number is multiplied by 10 to ensure

	var totalSum int

	// IMPORTANT: To pass parameter by value and not by reference
	var dfs func(node *TreeNode, currVal int)
	dfs = func(node *TreeNode, currVal int) {

		// When there is no tree left
		if node == nil {
			return
		}

		currVal = currVal*10 + node.Val

		// When we hit the leaf node
		if node.Left == nil && node.Right == nil {
			totalSum += currVal
			return
		}

		dfs(node.Left, currVal)

		dfs(node.Right, currVal)
	}

	dfs(root, 0)
	return totalSum
}
