package main

func cloneGraphBFS(node *Node) *Node {

	if node == nil {
		return nil
	}

	// Initialize the clone with the root of the original graph
	clone := &Node{Val: node.Val}

	// Make a visit map
	visit := make(map[*Node]*Node)

	// Map the original node to the clone node
	visit[node] = clone

	var queue []*Node
	queue = append(queue, node)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbour := range curr.Neighbors {
			if _, isVisited := visit[neighbour]; !isVisited {

				// Clone the neighbour
				cloneNeighbour := &Node{Val: neighbour.Val}

				// Mark the neighbour as visited
				visit[neighbour] = cloneNeighbour

				// Copy the original node neighbours to the clone node neighbours
				visit[curr].Neighbors = append(visit[curr].Neighbors, cloneNeighbour)

				// Enqueue the current neighbours from the original node who are not visited yet
				queue = append(queue, neighbour)
			} else {
				visit[curr].Neighbors = append(visit[curr].Neighbors, visit[neighbour])
			}
		}
	}

	return clone
}
