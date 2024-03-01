package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

	// Straight forward approach: Do an in-order traversal form the array and return arr[k-1].
	// SIMPLE BUT MORE EFFICIENT approach: Decrement k everytime you visit (in-order) the node and when k hits
	// zero we have the kth smallest element

	// result := inorderTraversal(root)
	// return result[k-1]

	var result int
	inorderTraversalTillK(root, &k, &result)
	return result
}

func inorderTraversalTillK(root *TreeNode, k, result *int) {

	// Inorder Traversal => left -> root -> right

	// Base case for recursion
	if root == nil {
		return
	}

	inorderTraversalTillK(root.Left, k, result)

	// Decrement the count until k hits zero
	*k--
	if *k == 0 {
		*result = root.Val
	}

	inorderTraversalTillK(root.Right, k, result)
}
