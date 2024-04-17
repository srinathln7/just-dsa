package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	// Key Idea: Perform a BFS for level-order traversal

	var result [][]int
	if root == nil {
		return result
	}

	queue := []*Node{root}

	var bfs func()
	bfs = func() {
		for len(queue) > 0 {
			var path []int
			levelSize := len(queue)

			for i := 0; i < levelSize; i++ {
				curr := queue[0]
				path = append(path, curr.Val)
				queue = queue[1:]
				queue = append(queue, curr.Children...)
			}

			result = append(result, path)
		}
	}

	bfs()
	return result
}
