# Check if Prerequisite

Given `n` courses and a list of prerequisites, represented as pairs of integers where each pair `[a, b]` represents that course `a` is a prerequisite of course `b`, and a list of queries where each query is represented as a pair `[a, b]` where `a` is the course and `b` is the target course. Return a list of boolean values where each element indicates whether course `a` is a prerequisite of course `b`.

## Approach

1. Create a 2D boolean array `connected` of size `n x n` to represent the connectivity between courses. Initialize all values to `false`.
2. For each prerequisite pair `[a, b]`, set `connected[a][b]` to `true` to mark course `a` as a prerequisite of course `b`.
3. Compute the transitive closure of the `connected` array using the Floyd-Warshall algorithm.
4. For each query `[a, b]`, check if `connected[a][b]` is `true` and store the result in the answer list.

## Time Complexity

Let `m` be the number of prerequisites and `q` be the number of queries.
- Constructing the connectivity matrix: O(m)
- Computing the transitive closure: O(n^3)
- Answering each query: O(q)
Overall time complexity: O(n^3 + m + q)

## Space Complexity

- Space required for the connectivity matrix: O(n^2)
- Space required for the answer list: O(q)
Overall space complexity: O(n^2 + q)
