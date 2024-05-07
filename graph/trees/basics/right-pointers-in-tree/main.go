package main

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	// Key Idea: BFS - Level order traversal . Form the queue slice at every iteration and connect the Next pointer until the end
	// of the slice

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize-1; i++ {
			queue[i].Next = queue[i+1]
		}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
