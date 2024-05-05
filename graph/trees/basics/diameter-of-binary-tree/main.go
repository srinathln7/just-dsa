package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	// Key Idea: Run a DFS on each node and keep track of the max. path length passing through that node
	// which the sum of the depth of the left and right sub-trees.

	// Run a DFS on each node and keep track of the max. path length passing through that node
	// which the sum of the depth of the left and right sub-trees. We do this for all nodes in the tree
	// and return the max path length i.e. diameter

	// IMPORTANT: Cause of confusion: The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
	// THIS PATH MAY or MAY NOT PASS THROUGH THE ROOT.
	// The dfs returns the depth of the node and NOT the depth of the tree (which is the max. number of nodes from ROOT to the LEAF).

	var diameter int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Run DFS with `node.Left` and `node.Right` as root
		leftPath, rightPath := dfs(node.Left), dfs(node.Right)
		diameter = max(diameter, leftPath+rightPath)
		return depth(node)
	}

	dfs(root)
	return diameter
}

func depth(node *TreeNode) int {

	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	return 1 + max(depth(node.Left), depth(node.Right))
}
