package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

	// Key Idea: We solve this problem using recursion by solving independent sub-problems. Notice that we cannot
	// flip the root itself but only its left and right sub trees. For the entire tree to be flip equivalent
	// its left and right subtrees need tobe flip equivalent

	// Base cases : Two empty trees are flip equivalent
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	return ((flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)))
}
