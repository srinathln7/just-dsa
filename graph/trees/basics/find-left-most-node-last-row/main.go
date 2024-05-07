package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {

	// Key Idea: To find the bottom most left value we cal. depth of both left and right subtrees and then narrow down the
	// route accorinding to whichever subtrees having the max. depth

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right)
	if rightDepth > leftDepth {
		return findBottomLeftValue(root.Right)
	}

	return findBottomLeftValue(root.Left)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth, rightDepth := maxDepth(root.Left), maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}
