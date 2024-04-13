package main

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	// Key Idea: Since we need to find the optimium (minimum) cost route with ATMOST `k` stops, we cannot apply Dijkstras algorithm straight forwardedly here.
	// An alternatvie algorithm that can help with constraints such as atmost `k` stops or nrgative weights in a directed graph is the Bellman-Ford algorithm.

	// Define a prices slice that represents the price to go from `src` city to any other cities.
	// Cities are zero-indexed and init it to pos. inf.
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = math.MaxInt32
	}

	// IMPORTANT: price to reach the `src` city from `src` city is always zero i.e. from = to
	prices[src] = 0

	// Since we are allowed atmost `k` stops, we can run the loop `k+1` times i.e. i -> [0..k]
	// Iteration `i` represents `i`th stop
	for i := 0; i <= k; i++ {

		// Create a temp slice to store the COPY of prices info. in the current iteration
		// temp := prices  => MISTAKE: Since the underlying price ref. are stored
		temp := make([]int, n)
		copy(temp, prices)

		for _, flight := range flights {
			// `s` -> source , `d` -> destination, `p` -> price
			s, d, p := flight[0], flight[1], flight[2]
			temp[d] = min(temp[d], prices[s]+p)

			// MISTAKE: We cannot simulataneously update the price of the destination with `temp[s]` since that would discard counting the intermediate stops
			// This is equivalent to not creating a `temp` copy of the prices slice and operating directly on it and running plain Dijkstra's algorithm.
			// For eg: Flights : [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], n=3, src=0, dst=3, k=1 => min.cost is 700 and not 400.
			// Dijkstras => Shortest path is 400 (0->1->2->3) involves `2` intermediate steps but here `k=1` (0->1->3).
			// temp[d] = min(temp[d], temp[s]+p)
		}

		prices = temp

	}

	// Despite running Bellman-Ford, if the prices have not changed from po. inf, it means there is no path from `src` to `dest`
	if prices[dst] == math.MaxInt32 {
		return -1
	}

	return prices[dst]
}
