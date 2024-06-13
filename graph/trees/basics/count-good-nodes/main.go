package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	// Key Idea: Perform a DFS and keep track of the maximum in While traversing each node, we keep track of the maximum value encountered from the root to the current node.
	// If the value of the current node is greater than or equal to this maximum value, it is considered a good node. By recursively traversing the tree and updating the maximum
	// value, we can count the number of good nodes in the binary tree.

	if root == nil {
		return 0
	}

	count := 0
	var dfs func(node *TreeNode, currMax int)
	dfs = func(node *TreeNode, currMax int) {

		if node == nil {
			return
		}

		if node.Val >= currMax {
			currMax = node.Val
			count++
		}

		dfs(node.Left, currMax)
		dfs(node.Right, currMax)
	}

	// IMPORTANT: To pass the `currMax` in the recursive stack frame function so that it takes only the
	// maximum value along the path from root to node
	dfs(root, root.Val)
	return count
}
