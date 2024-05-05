package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {

	// Key Idea: Preorder Traversal => root -> left -> right. Recursively flatten the left subtree and then the right subtree
	// The main challenge is we need to stick the flattened left subtree inbetween the root and the right subtree and we do
	// this by finding the right most child of the left subtree and point to the right subtree.

	// Base cases:
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	// Temporarily store the right subtree
	temp := root.Right

	// Flatten the trees in pre-order fashion
	flatten(root.Left)

	flatten(root.Right)

	// The root's right pointer now points to the flatten left subtree
	root.Right = root.Left

	root.Left = nil

	// Find the right most child of the left subtree
	for root.Right != nil {
		root = root.Right
	}

	// Point the right most child of the left subtree to the original flattened right subtree
	root.Right = temp
}
