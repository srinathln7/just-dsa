package main

func DiameterOfBinaryTree(root *TreeNode) int {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	var diameter int
	dfs(root, &diameter)
	return diameter
}

func dfs(root *TreeNode, diameter *int) int {

	if root == nil {
		return 0
	}

	leftDepth := dfs(root.Left, diameter)
	rightDepth := dfs(root.Right, diameter)

	*diameter = max(*diameter, leftDepth+rightDepth)
	return 1 + max(leftDepth, rightDepth)
}
