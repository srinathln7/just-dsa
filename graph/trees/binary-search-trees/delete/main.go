package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	case key == root.Val:
		switch {
		case root.Left == nil:
			root = root.Right
		case root.Right == nil:
			root = root.Left
		// When both left and right sub-trees are present, we find the node with minimum value in the right-most subtree
		// and replace the deleted node with this node. By this way, the BST property left_sub_tree < root <right_sub_tree
		// will still be PRESERVED.
		default:
			minNode := findMin(root.Right) // // Note pointer passed as value and not as reference
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
