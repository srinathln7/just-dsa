package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// To account for nodes that has only negative values
	maxSum := math.MinInt32
	dfs(root, &maxSum)
	return maxSum
}

func dfs(curr *TreeNode, maxSum *int) int {
	if curr == nil {
		return 0
	}

	// Calculate the maximum path sum in the left and right subtrees
	// If the sum is negative in either of the sub-trees, then there is no use going down that path
	// since we are trying to find the max path sum value
	leftSum := max(0, dfs(curr.Left, maxSum))
	rightSum := max(0, dfs(curr.Right, maxSum))

	// Update the maximum path sum found so far
	*maxSum = max(*maxSum, leftSum+rightSum+curr.Val)

	// Return the maximum path sum starting from the current node
	return max(leftSum, rightSum) + curr.Val
}
