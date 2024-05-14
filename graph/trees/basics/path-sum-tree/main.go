package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	// Key Idea: Use recursion and every time you encounter the node subract its value from target value and then finally check
	// if the leaf node's value equals the target value

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
