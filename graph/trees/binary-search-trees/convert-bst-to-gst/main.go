package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {

	// Key Idea: Reverse inorder traversal => Right sub_tree -> Root -> Left sub_tree

	if root == nil {
		return nil
	}

	// IMPORTANT MISTAKE: var revInorderTraversal func(node *TreeNode, val int) => Passing value by value and not reference

	// The bug in the provided code lies in the way the val variable is being updated and passed down during the reverse inorder traversal.
	// Since Go passes arguments by value, the updates made to val inside the recursive function revInorderTraversal are not reflected in the parent call
	// or in the next recursive call. To fix this issue, we need to pass the val variable by reference so that changes made to it inside the recursive function
	// are reflected in the outer scope. We can achieve this by passing a pointer to val instead of its value.

	var revInorderTraversal func(node *TreeNode, val *int)
	revInorderTraversal = func(node *TreeNode, val *int) {

		if node == nil {
			return
		}

		revInorderTraversal(node.Right, val)

		node.Val += *val
		*val = node.Val

		revInorderTraversal(node.Left, val)
	}

	sum := 0
	revInorderTraversal(root, &sum)
	return root

}
