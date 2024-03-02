package main

func cloneGraphDFS(node *Node) *Node {

	if node == nil {
		return nil
	}

	// Map to store the mapping between original nodes and their clones
	visited := make(map[*Node]*Node)

	return cloneDFS(node, visited)
}

func cloneDFS(node *Node, visited map[*Node]*Node) *Node {

	// If the node has already been visited, return its clone
	if clone, isVisited := visited[node]; isVisited {
		return clone
	}

	// Create the clone of the current node
	clone := &Node{Val: node.Val}

	// Mark the node as visited
	visited[node] = clone

	// Recursively clone the neighbors of the current node
	for _, neigbour := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneDFS(neigbour, visited))
	}

	return clone
}
