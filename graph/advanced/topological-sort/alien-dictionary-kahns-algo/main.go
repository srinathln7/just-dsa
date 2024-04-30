package main

import "fmt"

func alienOrder(words []string) string {

	// Key Idea: We run topological sort using Kahn's algorithm by calculating the in-degree of every vertex
	// We run BFS by populating the queue according to the in-degree of the vertices

	adj := make(map[byte][]byte)
	inDegree := make(map[byte]int)

	// Init adj map and indegree
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			adj[word[i]] = []byte{}
			inDegree[word[i]] = 0
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))

		// Words with same prefix but invalid ordering of length => return empty string
		if w1[:minLen] == w2[:minLen] && len(w1) > len(w2) {
			return ""
		}

		for j := 0; j < minLen; j++ {

			// Ordering can be determined through the first unequal characters in both the words
			if w1[j] != w2[j] {
				adj[w1[j]] = append(adj[w1[j]], w2[j])
				inDegree[w2[j]]++
				break
			}
		}
	}

	// Topological sort using Kahn's algorithm
	queue := make([]byte, 0)
	for ch := range adj {
		if inDegree[ch] == 0 {
			queue = append(queue, ch)
		}
	}

	var order []byte
	for len(queue) > 0 {
		node := queue[0]
		order = append(order, node)

		queue = queue[1:]

		for _, neighbour := range adj[node] {
			inDegree[neighbour]--

			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}
		}
	}

	// Check for cycle
	if len(order) != len(adj) {
		return ""
	}

	return string(order)
}

func main() {
	words0 := []string{"wrt", "wrf", "er", "ett", "rftt"}
	words1 := []string{"z", "x"}
	words2 := []string{"z", "x", "z"}

	order0 := alienOrder(words0)
	order1 := alienOrder(words1)
	order2 := alienOrder(words2)

	fmt.Println("Alien dictionary order:", order0)
	fmt.Println("Alien dictionary order:", order1)
	fmt.Println("Alien dictionary order:", order2)
}
