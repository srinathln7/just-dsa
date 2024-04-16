package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	// Key Idea: We see the newly inverted tree is the mirror image of the actual tree.
	// Let us recursively invert the tree and all its sub-trees.

	// Base cases: If there is no tree
	if root == nil {
		return root
	}

	// When we reach the leaf node
	if root.Left == nil && root.Right == nil {
		return root
	}
	var newroot *TreeNode = &TreeNode{Val: root.Val}
	newroot.Left = invertTree(root.Right)
	newroot.Right = invertTree(root.Left)
	return newroot
}
