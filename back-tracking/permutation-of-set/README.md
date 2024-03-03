# Permutation of a set

Given an array `nums` of distinct integers, return all possible permutations. You can return the answer in any order.

## Intution 
The problem can be solved using Depth First Search (DFS) with backtracking. The approach involves exploring all possible permutations by recursively choosing each element from the input array and tracking which elements have been used.

## Approach

1. Implement a `permute` function to initialize the result and start the DFS.
2. Implement a `dfs` function that performs the depth-first search.
3. In the `dfs` function, if the size of the permutation equals the length of the input array, append the permutation to the result.
4. Otherwise, iterate through each element in the input array (`nums`) and check if it has been used.
5. If the element has not been used, include it in the current permutation, mark it as used, and recursively call the `dfs` function.
6. After the recursive call, backtrack by removing the last element from the permutation and marking the element as unused.
7. Repeat steps 4-6 for each element in the input array to explore all possible permutations.

## Time Complexity

The time complexity of this approach is O(N!), where N is the length of the input array. This is because there are N choices for the first element, N-1 choices for the second element, and so on, resulting in N! possible permutations.

## Space Complexity

The space complexity is O(N), where N is the length of the input array. This is because we use additional space for the result, used map, and the current permutation during the DFS traversal. However, the recursion stack also contributes to space usage, but it is limited to O(N) in this case.
