package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	// Key Idea: Let us recursively build the left and right sub-trees. Notice that the `mid` will form the root of the tree.
	// All elements to the left of `mid` will form the left sub tree and all elements to the right of `mid` will form the
	// right sub tree.

	n := len(nums)

	// Base cases
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	l, r := 0, n-1
	mid := l + (r-l)/2

	newRoot := &TreeNode{Val: nums[mid]}
	newRoot.Left, newRoot.Right = sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])
	return newRoot
}
