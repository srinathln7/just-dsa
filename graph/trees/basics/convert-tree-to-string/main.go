package main

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {

	// Key Idea: Do a pre-order DFS traversal passing the result as the reference to the string
	// This is because the result variable needs tobe mutated across the recursive stack frame

	if root == nil {
		return ""
	}

	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}

	var dfs func(node *TreeNode, result *string)
	dfs = func(node *TreeNode, result *string) {
		if node == nil {
			return
		}

		// Pre-order Trav => root -> left sub_tree -> right sub_tree
		*result += strconv.Itoa(node.Val)

		if node.Left == nil && node.Right == nil {
			return
		}

		// Empty paranthesis should be present to indicate the absence of a left child
		*result += "("
		dfs(node.Left, result)
		*result += ")"

		// Empty paranthesis needs to be omitted in case there is no right child
		if node.Right != nil {
			*result += "("
			dfs(node.Right, result)
			*result += ")"
		}

	}

	var result string
	dfs(root, &result)

	return result
}
