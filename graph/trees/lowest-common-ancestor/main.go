package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// Key Idea: BST => All vals in Left sub_tree < Root vals < Right sub_tree
	// So, we know if either p's val is lesser than or equal to the root val and q's val is greater than or equal to root's val
	// or vice-versa the lowest common ancestor (LCA) is always the root. If both the values are strictly lesser than the root's val
	// the LCA lies to the left of the sub-tree. We can find this recursively. Otherwise, the LCA lies to the right of the sub-tree

	if (p.Val <= root.Val && root.Val <= q.Val) || (q.Val <= root.Val && root.Val <= p.Val) {
		return root
	}

	// LCA on the left
	if root.Left != nil && (p.Val < root.Val && q.Val < root.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// LCA on the right
	return lowestCommonAncestor(root.Right, p, q)
}
