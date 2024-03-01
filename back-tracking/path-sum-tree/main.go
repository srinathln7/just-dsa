package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	// We can solve this problem with BACK TRACKING algorithm which in essence is a brute-force algorithm
	if root == nil {
		return false
	}

	// Traverse from the root all the way to the leaf

	// Leaf node condition
	if (root.Left == nil && root.Right == nil) && root.Val == targetSum {
		return true
	}

	// Check in the left subtree
	if root.Left != nil && hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}

	// Check in the right subtree
	if root.Right != nil && hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return false
}
