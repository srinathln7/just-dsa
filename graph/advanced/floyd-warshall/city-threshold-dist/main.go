package main

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	// Key Idea: Let us find the shortest path from all the vertex to all other vertices in the graph using
	// the Floyd-Warshall algorithm through the approach of bottom-up dynamic programming. Then let us loop
	// over the results to output the city with the greatest number that meets the `distanceThreshold`

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}

	// Init the `dist` (DP) 2D-array to pos. inf
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Distance from `src` to `src` is always zero
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	// Fill the DP array with the given weight of the edges
	for _, edge := range edges {
		src, dst, weight := edge[0], edge[1], edge[2]
		// Bi-directional graph
		dist[src][dst] = weight
		dist[dst][src] = weight
	}

	// Loop over every vertex in the graph as possible intermediate vertex `k`
	// to find the shortest possible path for all possible pairs of vertices
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	// Range over the filled-in DP array to get the result
	// We also keep track of the neighbours for our final result
	neighbours := make([]int, n)
	for src := 0; src < n; src++ {
		for dst := 0; dst < n; dst++ {
			if src != dst && dist[src][dst] <= distanceThreshold {
				neighbours[src]++
			}
		}
	}

	// Largest city With the smallest number of neighbors at a threshold distance
	result, min := 0, neighbours[0]
	for i := 1; i < n; i++ {
		if neighbours[i] <= min {
			min = neighbours[i]
			result = i
		}
	}

	return result
}
