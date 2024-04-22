package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	// Store mapping from original node to copy node
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(curr *Node) *Node {
		if copyNode, exists := visited[curr]; exists {
			return copyNode
		}

		copyNode := &Node{Val: curr.Val}

		// Mark the node as visited
		visited[curr] = copyNode
		for _, neighbor := range curr.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighbor))
		}

		return copyNode
	}

	return dfs(node)
}
