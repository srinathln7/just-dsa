package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	// Key Idea: Run a DFS on each node and keep track of the max. path length passing through that node
	// which the sum of the depth of the left and right sub-trees. We do this for all nodes in the tree
	// and return the max path length i.e. diameter

	// IMPORTANT: Cause of confusion: The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
	// THIS PATH MAY or MAY NOT PASS THROUGH THE ROOT.
	// The dfs returns the depth of the node and NOT the depth of the tree (which is the max. number of nodes from ROOT to the LEAF).

	res := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res = max(res, left+right)
		return 1 + max(left, right)
	}

	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
