# Topological Ordering
Given a list of edges representing a directed acyclic graph (DAG), the task is to find a valid topological ordering of the graph.

## Intuition
Topological sorting is a way of arranging the nodes of a directed graph in a linear ordering such that for every directed edge from node u to node v, u comes before v in the ordering. This ordering is not unique, and there may be multiple valid topological orderings for a given DAG.

## Approach
1. **Forming Adjacency Graph**: First, the code forms the adjacency graph from the given list of edges. It uses a map where the keys represent the nodes, and the values represent the list of adjacent nodes.
2. **DFS Traversal**: Next, it performs a depth-first search (DFS) traversal on the graph. During the traversal, it marks visited nodes and recursively explores the neighboring nodes.
3. **Building Topological Order**: As the DFS traversal progresses, it adds the nodes to the topological ordering in reverse order.
4. **Reverse Order**: Finally, once the full topological ordering is built, it reverses the order to obtain the correct topological sorting.

## Time Complexity
The time complexity of this algorithm is O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph. This complexity arises due to the DFS traversal.

## Space Complexity
The space complexity of this algorithm is O(V + E), where V is the number of vertices and E is the number of edges. This space is used to store the adjacency graph, visited nodes, and the resulting topological ordering.

## Example
For the given example with edges `[[1, 3], [1, 4], [2, 4], [3, 5], [4, 6], [5, 6]]` and `n = 6`, the output topological sorting is `[2 1 4 3 5 6]`.

Ordered: Top-down i.e. arrow points downwards
``` 
      1  2
     / \/
    3   4 
   /     \
  5       6
 
```

The code successfully implements topological sorting using depth-first search on a directed acyclic graph.