package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	// Key Idea: Let us recursively merge the two roots and its sub trees.

	// Base cases - When either one of the tree is empty
	if root2 == nil {
		return root1
	}

	if root1 == nil {
		return root2
	}

	// Leaf nodes
	if isLeafNode(root1) && isLeafNode(root2) {
		return &TreeNode{Val: root1.Val + root2.Val}
	}

	newRoot := &TreeNode{Val: root1.Val + root2.Val}
	newRoot.Left = mergeTrees(root1.Left, root2.Left)
	newRoot.Right = mergeTrees(root1.Right, root2.Right)

	return newRoot
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
