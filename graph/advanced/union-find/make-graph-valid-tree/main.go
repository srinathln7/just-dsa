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
	// Using `uf.parent[n1]` and `uf.parent[n2]` directly to find the parents of the nodes instead of the `find()` algorithm is incorrect because
	// it does not take into account path compression. In a proper implementation of a Union-Find data structure, the `find()` method not only finds
	// the parent of a node but also compresses the path to the root. This path compression step optimizes subsequent `find()` operations, making them more efficient.
	// Without path compression, the depth of the tree formed by the union-find operations could become linear in the worst case, leading to inefficient performance,
	// especially for large datasets or complex graphs. Therefore, it's essential to use the `find()` method to ensure proper path compression and maintain the
	// efficiency of the Union-Find data structure.
	p1, p2 := uf.find(n1), uf.find(n2)

	//p1, p2 := uf.parent[n1], uf.parent[n2] => MISTAKE

	// Common ancestor - Cycle detected
	if p1 == p2 {
		return false
	}

	// Perform union operation based on the rank. Node with the highest rank becomes a parent to the node with the lowest rank. This is done to balance the tree.
	// IMPORTANT to increment the rank only when the two node levels are the same. Otherwise, the node with greater rank becomes the parent and hence
	// the height (rank) does not increase/decrease as such
	switch {
	case uf.rank[p1] > uf.rank[p2]:
		uf.parent[p2] = p1
	case uf.rank[p1] < uf.rank[p2]:
		uf.parent[p1] = p2
	default:
		uf.parent[p2] = p1
		uf.rank[p1]++
	}

	return true
}

func findRedundantConnection(edges [][]int) []int {
	// Create Union find over the graph
	n := len(edges)
	uf := newUnionFind(n)

	// Since, we have to return the last redundant connection incase of multiple
	// redundant connections, we over the result with the edge after every iteration.
	var result []int
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		// Since union find oeprates only on dis-joint sets => if union operation fails,
		// it indicates the redundant connection or presence of a cycle i.e. the two nodes are
		// already part of the same sub-graph. ALTERNATIVELY, if we are
		// asked to find the number of connected components, we can get it by calculating
		// `#nodes - #unionfind`
		if !uf.union(n1, n2) {
			result = edge
		}
	}

	return result
}
