package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	// Inorder Traversal => left_subtree -> root -> right_subtree
	var result []int
	switch {
	case root == nil:
	case root.Left == nil && root.Right == nil:
		result = append(result, root.Val)
	default:
		leftSubTree := inorderTraversal(root.Left)
		rightSubTree := inorderTraversal(root.Right)
		result = leftSubTree
		result = append(result, root.Val)
		result = append(result, rightSubTree...)
	}
	return result
}
