# Alien Dictionary 
Given a list of strings `words` representing an alien dictionary, where the letters in the dictionary are ordered lexicographically, determine the order of the characters in the alien language. If the order is invalid or ambiguous, return an empty string.

## Intuition
To determine the order of characters in the alien language, we can consider the given list of words as relationships between characters. We need to find a valid ordering of characters such that all given words are sorted lexicographically. This problem can be solved by applying topological sort using Kahn's algorithm.

## Approach
1. Initialize an adjacency list `adj` to represent the relationships between characters and a map `inDegree` to store the in-degree of each character.
2. Iterate through the list of words to construct the adjacency list and calculate the in-degree of each character.
3. For each pair of adjacent words, find the first differing characters. The character in the first word comes before the character in the second word.
4. Perform topological sort using Kahn's algorithm:
   - Initialize a queue and enqueue all characters with an in-degree of 0.
   - While the queue is not empty, dequeue a character `node`:
     - Append `node` to the result order.
     - Update the in-degree of its neighbors and enqueue any neighbor with an in-degree of 0.
   - If the length of the result order is not equal to the number of characters, there is a cycle in the graph, so return an empty string.
5. Return the result order as a string.

## Time Complexity
- Constructing the adjacency list and calculating in-degrees: O(N), where N is the total number of characters in all words.
- Topological sort using Kahn's algorithm: O(V + E), where V is the number of vertices (characters) and E is the number of edges (relationships between characters) in the adjacency list.
- Overall time complexity: O(N + V + E)

## Space Complexity
- Space required for the adjacency list and in-degree map: O(V + E)
- Space required for the queue: O(V)
- Overall space complexity: O(N + V + E)