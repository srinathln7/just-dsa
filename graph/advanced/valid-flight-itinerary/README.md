# Valid Itinerary

You are given a list of airline tickets represented by pairs of departure and arrival airports `[from, to]`, where each element `tickets[i] = [fromi, toi]` represents a ticket from `fromi` to `toi`. The goal is to reconstruct the itinerary in order such that the tickets are used in the correct order to form a complete trip. It is guaranteed that the input represents a valid itinerary. More specifically, the input must form a circuit where each ticket is used once and only once.

## Intuition

To reconstruct the itinerary, we need to find the path that starts from the departure airport `"JFK"` and follows the smallest lexical order of destination airports. We can represent the given tickets as a graph using an adjacency list. By sorting the destinations in lexical order for each departure airport, we can ensure that we always choose the smallest lexical order destination first when traversing the graph. We then perform a depth-first search (DFS) starting from the airport `"JFK"` to construct the itinerary. If the length of the resulting itinerary is equal to the number of tickets plus one, we have found a valid itinerary.

## Approach

1. Form an adjacency list representing the graph using the given tickets. Each key in the map represents a departure airport, and the corresponding value is a list of arrival airports.
2. Sort the adjacency list in lexical order for each departure airport.
3. Implement a depth-first search (DFS) function that explores the graph recursively, starting from the airport `"JFK"`. During DFS traversal, always choose the smallest lexical order destination first.
4. After DFS traversal, check if the length of the resulting itinerary is equal to the number of tickets plus one. If so, return the itinerary; otherwise, return nil to indicate that no valid itinerary exists.
5. Reverse the resulting itinerary to ensure the correct order.

## Time Complexity

For a typical DFS O(V+E) since we have to visit every vertex and edge atleast once. However, since backtracking is also involved here for every possible edge,
O((V+E). (V+E)) ~ O(E^2) since (E <= V^2)

### Space Complexity

The overall space complexity is O(E) where `E` is the number of edges (tickets).