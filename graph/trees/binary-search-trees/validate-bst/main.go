package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// BST => left subtree < root < right subtree
	// We introduce the concept of min and max tree to check if a given tree is a BST
	return isvalidBST(root, nil, nil)
}

func isvalidBST(root, min, max *TreeNode) bool {
	// Empty tree is a valid BST
	if root == nil {
		return true
	}

	if (max != nil && root.Val >= max.Val) || (min != nil && root.Val <= min.Val) {
		return false
	}

	// Recursively check for the BST property in left and right subtree
	// Max value of a left subtree must be the ROOT and Min value of the right subtree must be the ROOT
	// IMPORTANT MISTAKE: make sure not to pass `return isvalidBST(root.Left, nil, root) && isvalidBST(root.Right, root, nil)`
	// as we need to check for all subtrees in left and right
	return isvalidBST(root.Left, min, root) && isvalidBST(root.Right, root, max)
}
