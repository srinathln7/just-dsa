package main

// UnionFind: A data structure that lets us efficiently find if a
// given undirected (bi-directional) graph has a cycle or not.
type UnionFind struct {
	parent map[int]int
	rank   map[int]int
}

// newUnionFind: Returns a new union find datastructure
func newUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}

	// Default values for union find: Parent pointing to itself and rank (height) set to zero
	for i := 1; i <= n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}

	return uf
}

// find: Returns the ancestor (parent/grand parent) for the given node.
// It is IMPORTANT to note that union find does not preserve the actual
// structure of the graph but is trying to find only the ancestor up to the
// root quickly.
func (uf *UnionFind) find(node int) int {
	p := uf.parent[node]
	for p != uf.parent[p] {
		// Find the grand parent of the parent
		// Run the PATH COMPRESSION algorithm to find the common ancestor upto the root
		// This also balances the tree
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p] // Update the parent to the grand parent
	}

	return p
}

// Union operates only on dis-joint sets i.e. it unites only when two sub-graphs (FOREST OF TREES) are
// not already connected to each other. If two nodes have a common ancestor it means they are a part of the same
// sub-graph and are already connected. We cannot connect two nodes that are already connected
// and hence returns false. Otherwise, it performs a union operation meaning it connectes
// the two nodes and form a single sub-graph
func (uf *UnionFind) union(n1, n2 int) bool {

	// Extract the parents of the two nodes using the union find algorithm

	//p1, p2 := uf.parent[n1], uf.parent[n2]
	p1, p2 := uf.find(n1), uf.find(n2)

	// Common ancestor - Cycle detected
	if p1 == p2 {
		return false
	}

	// Perform union operation based on the rank. Node with the
	// highest rank becomes a parent to the node with the lowest rank
	// This is done to balance the tree
	switch {
	case uf.rank[p1] > uf.rank[p2]:
		uf.parent[p2] = p1
	case uf.rank[p1] < uf.rank[p2]:
		uf.parent[p1] = p2
	default:
		uf.parent[p2] = p1
		uf.rank[p1]++ // IMPORTANT to increment the rank only when the two node levels are the same
	}

	return true
}

func findRedundantConnection(edges [][]int) []int {
	// Create Union find over the graph
	n := len(edges)
	uf := newUnionFind(n)

	// Since, we have to return the last redundant connection incase of multiple
	// redundant connections, we overwrite the result with the edge after every iteration.
	var result []int
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		// If union operation fails, it indicates the presence of a cycle i.e. the two nodes are
		// already part of the same sub-graph i.e. a redundant connection
		if !uf.union(n1, n2) {
			result = edge
		}
	}

	return result
}
