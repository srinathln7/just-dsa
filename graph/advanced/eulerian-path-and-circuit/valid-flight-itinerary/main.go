package main

import "sort"

func findItinerary(tickets [][]string) []string {

	// Key Idea: To form the graph using adjancency list and run a DFS starting from node "JFK". Since, we have to return the
	// itinerary with the smallest lexical order, we first sort the adj list. As we progress along DFS from the `src` point,
	// we backtrack if there is no further path. At the end of the recursion, we append the result. Since we have use every ticket
	// once and only once and we start from a src node `JFK` (+1), there exists a valid itinerary if the len(result) == len(tickets) + 1

	// Form the adjancey list
	adj := make(map[string][]string)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		adj[src] = append(adj[src], dst)
	}

	// Sort the adj as per lexical order
	for src := range adj {
		sort.Strings(adj[src])
	}

	var result []string
	var dfs func(src string)
	dfs = func(src string) {
		for len(adj[src]) > 0 {
			dst := adj[src][0]

			// Since we can use the ticket only once, remove that city in the adj list of the src
			// before running the DFS
			adj[src] = adj[src][1:]

			// Run DFS on the neighbours as per the lexical order
			dfs(dst)
		}

		result = append(result, src)
	}

	// Perform a depth-first search (DFS) or Hierholzer's algorithm to find the Eulerian path
	// Run the DFS with `JFK` as the src node
	dfs("JFK")

	// Check if a valid itinerary exists
	if len(result) != len(tickets)+1 {
		return nil
	}

	// Since the result from the DFS uses a recursive solution, it appends the src node at the last
	// To output the result in valid order, we need to reverse
	rev(result)
	return result
}

func rev(str []string) {
	n := len(str)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
