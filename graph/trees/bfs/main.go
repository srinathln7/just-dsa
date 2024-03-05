package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	// Standard BFS-traversal using a dequeue
	var result [][]int
	if root == nil {
		return result
	}

	// Use a double-ended queue (aka) deque
	deque := []*TreeNode{root}
	var levelSize int
	for len(deque) > 0 {
		levelSize = len(deque)

		var level []int // Elements stored at every level
		for i := 0; i < levelSize; i++ {
			node := deque[0]
			level = append(level, node.Val)

			deque = deque[1:]

			if node.Left != nil {
				deque = append(deque, node.Left)
			}

			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
