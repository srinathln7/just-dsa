package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	// Key Idea: We recursively calculate the depth of both the left and right sub-trees and take the maximum of both

	// Base case

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)
	return max(leftDepth, rightDepth)
}
