package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return &TreeNode{Val: val}
	case val < root.Val:
		root.Left = insertIntoBST(root.Left, val) // Note we are updating the `root.Left` pointer in this case
	case val > root.Val:
		root.Right = insertIntoBST(root.Right, val) // Note we are updating the `root.Right` pointer in this case
	}

	return root
}
