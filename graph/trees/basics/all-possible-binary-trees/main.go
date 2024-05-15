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

	// dp[i] => No. of full binary tree with `i` number of nodes

	dp := make([][]*TreeNode, n+1)
	node := &TreeNode{}
	dp[1] = []*TreeNode{node}

	for i := 3; i <= n; i += 2 {
		// Total of `i-1` nodes
		for j := 1; j < i; j++ {
			leftTrees := dp[j-1]
			rightTrees := dp[i-j]
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					dp[i] = append(dp[i], &TreeNode{Left: left, Right: right})
				}
			}
		}
	}
	return dp[n]
}
