# Fixed (0/1) Knapsack 

Given a set of items, each with a weight and a profit, determine the maximum profit you can obtain by selecting a subset of the items such that the sum of the weights of the selected items is less than or equal to a given capacity.


## Bottom-up DP


### Intuition

We can solve this problem using dynamic programming. We visualize a 2D grid where rows represent items and columns represent capacity. We fill up this grid item-by-item and capacity-by-capacity. We can calculate the maximum profit for each cell by considering two decisions: including or excluding the current item. We optimize the space complexity from O(n*m) to O(m) by observing that the max-profit in the current cell depends only on the previous row.

### Approach

1. Initialize a 1D array `dp` of size `capacity + 1` to store the maximum profit for each capacity.
2. Fill up the first row of the `dp` array with maximum profits achievable by considering only the first item.
3. Iterate through the items and for each item, iterate through the capacity range.
4. For each capacity, calculate the maximum profit by considering two decisions:
    - Skip (exclude) the current item: profit remains the same as adding only the previous item.
    - Include the current item: calculate profit by adding the current item's profit with the profit obtained from adding the previous item with remaining capacity.
5. Update the `dp` array with the maximum profit for each capacity.
6. Return the maximum profit stored in the last cell of the `dp` array.

### Time Complexity

The time complexity is O(n * m), where n is the number of items and m is the capacity. We iterate through all items and for each item, we iterate through all capacities.

### Space Complexity

The space complexity is O(m), where m is the capacity. We only use a 1D array to store the maximum profit for each capacity, optimizing the space from the traditional 2D grid approach.



## Brute-Force DFS

### Intuition

We can solve this problem using a brute-force approach known as Depth-First Search (DFS). We traverse through all possible combinations of items and choose the one with maximum profit while ensuring that the capacity constraint is satisfied.

### Approach

1. Implement a DFS function `dfs` that takes in the weights, profits, and capacity as input parameters.
2. Initialize a helper function `dfsHelper` that performs the recursive depth-first search.
3. In the `dfsHelper` function, handle the base case where all items have been considered (index equals the length of the profits array).
4. Formulate two decisions:
   - Decision 1: Skip the current item and move to the next item. Recur with the updated index.
   - Decision 2: Include the current item if its weight can fit into the remaining capacity. Recur with the updated capacity and index.
5. Return the maximum profit obtained from the two decisions.
6. Call the `dfs` function with the initial parameters to get the maximum profit.

### Time Complexity

The time complexity of this approach is O(2^n), where n is the number of items. This is because for each item, we have two choices: include it or exclude it.

### Space Complexity

The space complexity is O(n), where n is the number of items. This is due to the recursive stack space used during the depth-first search traversal.
