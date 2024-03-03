# Backtracking

Given an array of distinct integers `nums`, return all possible subsets (the power set). The solution should not contain any duplicate subsets.

## Approach

The problem can be solved using backtracking, a technique that systematically searches through all possible solutions to find the optimal one. In this approach, we recursively build subsets by including or excluding each element in the array.

1. Implement a backtracking function `dfs` that recursively generates subsets.
2. Initialize an empty `result` slice to store the subsets.
3. Iterate through the elements of the input array `nums`.
4. In each iteration, include the current element in the subset and recursively call `dfs` with the updated subset and the next index.
5. After the recursive call, backtrack by removing the last element from the subset to explore other possibilities.
6. Append the subset to the `result` slice after each recursive call.
7. Return the `result` slice containing all possible subsets.

## Time Complexity

The time complexity of the backtracking algorithm is O(2^N), where N is the number of elements in the input array `nums`. This is because each element can either be included or excluded in a subset, resulting in a total of 2^N possible subsets.

## Space Complexity

The space complexity is O(N * 2^N), where N is the number of elements in the input array `nums`. This space is required to store all possible subsets generated during the backtracking process.


## Remark

The line `subset = subset[:len(subset)-1]` is significant because it removes the last element from the subset. In the context of backtracking algorithms, this line is used to backtrack or undo the inclusion of the current element in the subset before exploring other possibilities.

Here's why this line is important:

1. **Backtracking**: After including an element in the subset and exploring all possible combinations with that element included, we need to backtrack to explore other combinations where that element is not included. Backtracking involves removing the last element added to the subset so that we can explore other branches of the recursion tree.

2. **Preparation for Next Iteration**: In the recursive function `backtrack`, after processing the current element and its subsets, we remove the last element from the `subset` before moving on to the next iteration. This prepares the `subset` for the next recursive call, ensuring that we correctly build subsets without the current element for the subsequent iterations.

3. **Maintaining State**: By removing the last element from the `subset`, we revert it to the state it was in before adding the current element. This ensures that we explore all possible combinations of elements without duplicates and without missing any combination.

In summary, the line `subset = subset[:len(subset)-1]` plays a crucial role in backtracking algorithms by allowing us to backtrack and explore different combinations of elements while constructing subsets.


## Output

```
go run main.go
result at i= 0 => [[]]
subset at i=0 => [1]
Recurse with start index=1
result at i= 1 => [[] [1]]
subset at i=1 => [1 2]
Recurse with start index=2
result at i= 2 => [[] [1] [1 2]]
subset at i=2 => [1 2 3]
Recurse with start index=3
result at i= 3 => [[] [1] [1 2] [1 2 3]]
Backtracked subset at i=2 => [1 2] 
Backtracked subset at i=1 => [1] 
subset at i=2 => [1 3]
Recurse with start index=3
result at i= 3 => [[] [1] [1 2] [1 2 3] [1 3]]
Backtracked subset at i=2 => [1] 
Backtracked subset at i=0 => [] 
subset at i=1 => [2]
Recurse with start index=2
result at i= 2 => [[] [1] [1 2] [1 2 3] [1 3] [2]]
subset at i=2 => [2 3]
Recurse with start index=3
result at i= 3 => [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3]]
Backtracked subset at i=2 => [2] 
Backtracked subset at i=1 => [] 
subset at i=2 => [3]
Recurse with start index=3
result at i= 3 => [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
Backtracked subset at i=2 => [] 
```
