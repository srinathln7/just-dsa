package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// BST => left sub_tree <= root < right sub_tree
	switch {
	case root == nil:
		return nil
	case val == root.Val:
		return root
	case val < root.Val:
		return searchBST(root.Left, val)
	case val > root.Val:
		return searchBST(root.Right, val)
	}

	return nil
}
