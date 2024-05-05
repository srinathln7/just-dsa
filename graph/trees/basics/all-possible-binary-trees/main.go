package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {

	// Key Idea: We use bottom-up dynamic programming approach and start simple. Let dp[i] represent the number of unique full binary
	// trees with `i` number of nodes

	// Base case: Impossible to have a full binary tree for even `n`
	if n%2 == 0 {
		return nil
	}

	dp := make([][]*TreeNode, n+1)

	node := &TreeNode{}
	dp[1] = []*TreeNode{node}

	for i := 3; i <= n; i += 2 {

		var trees []*TreeNode

		// IMPORTANT : j < i as 1 node is reserved for root

		for j := 1; j < i; j += 2 {

			// Get all possible left subtrees with j nodes
			leftSubTrees := dp[j]

			// Get all possible right subtrees with i-j-1 nodes
			rightSubTrees := dp[i-j-1]

			// Generate all combinations of left and right subtrees to form full binary trees
			for _, left := range leftSubTrees {
				for _, right := range rightSubTrees {
					tree := &TreeNode{Left: left, Right: right}
					trees = append(trees, tree)
				}
			}
		}
		// Store the list of full binary trees with `i` number of nodes
		dp[i] = trees
	}

	// Return the list of full binary trees with `n` number of nodes
	return dp[n]
}
