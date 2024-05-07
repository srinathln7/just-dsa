package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {

	// Key Idea: Narrow down according to the range [low, high].
	// Use BST property i.e. All vals in left sub_tree < root Val < All Vals in right subtree

	// Base case : Empty tree
	if root == nil {
		return nil
	}

	// If root's val itself exceeds upper bound - all vals to the right will naturally exceed
	// Hence no use triming right subtree
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	// If root's val itself subseds lower bound - all vals to the left will be naturally lower
	// Hence no use triming left subtree
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	// If the root. val is within the range then start trimming the subtrees recursively
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
