package main

func RightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var levelTraverse [][]int

	deque := []*TreeNode{root}

	var levelSize int
	var node *TreeNode
	for len(deque) > 0 {
		levelSize = len(deque)

		var level []int
		for i := 0; i < levelSize; i++ {
			node = deque[0]
			level = append(level, node.Val)

			deque = deque[1:]

			if node.Left != nil {
				deque = append(deque, node.Left)
			}

			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}

		levelTraverse = append(levelTraverse, level)
	}

	result := make([]int, len(levelTraverse))
	for i := 0; i < len(result); i++ {
		result[i] = levelTraverse[i][len(levelTraverse[i])-1]
	}

	return result
}
