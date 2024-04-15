Eulerian path and circuit 
Given a graph, determine if it has an Eulerian path or an Eulerian circuit, and if so, find it.

Ex: 

```
      0
     / \
    1   2
   / \ / \
  3   4   3

```

## Intuition:  
An Eulerian path in a graph is a path that visits every edge exactly once. An Eulerian circuit is an Eulerian path that starts and ends at the same vertex. The existence of an Eulerian path or circuit in a graph depends on the degrees of its vertices.  It's important to note that an Eulerian path can exist in a graph even if the graph does not have an Eulerian circuit.

## Approach:
1. **Eulerian Path:**  
   - If a graph has an Eulerian path, it means it has exactly two vertices with odd degree, and all other vertices have even degree.
   - To find the Eulerian path, we need to start at one of the vertices with odd degree and traverse each edge exactly once, until we reach the other vertex with odd degree.

2. **Eulerian Circuit:**  
   - If a graph has an Eulerian circuit, it means all vertices have even degree.
   - To find the Eulerian circuit, we start at any vertex and traverse each edge exactly once, until we return to the starting vertex.

## Time and Space Complexity:  
- The time complexity of both algorithms is O(V + E), where V is the number of vertices and E is the number of edges in the graph.
- The space complexity is also O(V + E) for storing the graph.



