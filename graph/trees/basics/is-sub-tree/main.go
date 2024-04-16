package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	// Key Idea: Check if the given two trees are the same. If not, then check if the left-subtree or the right-sub tree
	// is same as the given sub-tree. If none of them is true, then return false.

	if isSameTree(root, subRoot) {
		return true
	}

	if (root.Left != nil && isSubtree(root.Left, subRoot)) || (root.Right != nil && isSubtree(root.Right, subRoot)) {
		return true
	}

	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// Key Idea: Lets recursively find if all the left-sub trees and right sub-trees are the same apart from checking the root values are the same

	// Base case: Two empty trees are the same identical trees
	if p == nil && q == nil {
		return true
	}

	// When one tree is empty and other one is not, then the two trees cannot be the same
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// Check if the root's value is equal and if the left and right sub-trees are the same
	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
}
