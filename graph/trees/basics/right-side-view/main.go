package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	deque := []*TreeNode{root}

	var levelSize int
	var node *TreeNode
	for len(deque) > 0 {
		levelSize = len(deque)

		// Key Idea: We start enqueuing first the root, then left and then the RIGHT
		// This means at every level, if an element exists in the right subtree, it gets
		// captured by the `rightSide` variable BEFORE the inner loop for every level exists.
		// Alt. if asked for left side view, first enque the right side element and then the left element
		var rightSide int
		for i := 0; i < levelSize; i++ {
			node = deque[0]
			rightSide = node.Val

			deque = deque[1:]

			if node.Left != nil {
				deque = append(deque, node.Left)
			}

			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}
		result = append(result, rightSide)
	}

	return result
}
