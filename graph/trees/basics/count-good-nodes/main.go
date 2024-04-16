package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	// Key Idea: Perform a DFS and keep track of the maximum in While traversing each node, we keep track of the maximum value encountered from the root to the current node.
	// If the value of the current node is greater than or equal to this maximum value, it is considered a good node. By recursively traversing the tree and updating the
	// maximum value, we can count the number of good nodes in the binary tree.

	var count int
	dfs(root, root.Val, &count)
	return count
}

func dfs(node *TreeNode, currMax int, count *int) {

	if node == nil {
		return
	}

	if node.Val >= currMax {
		currMax = node.Val
		*count++
	}

	dfs(node.Left, currMax, count)
	dfs(node.Right, currMax, count)
}
