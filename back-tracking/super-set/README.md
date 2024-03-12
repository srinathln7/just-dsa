# Superset
Given a set of distinct integers, nums, return all possible subsets (the power set).

## Intuition
To generate all possible subsets, we can use a depth-first search (DFS) approach. At each step, we have two choices: include the current element in the subset or exclude it. We recursively explore both options until we reach the end of the array.

## Approach
1. Define a function `dfs` to perform the depth-first search. The function takes an integer `i` representing the current index in the array.
2. If `i` equals the length of the array `nums`, it means we have processed all elements, so we append a copy of the current subset to the result.
3. Otherwise, we have two choices:
   - Include the current element `nums[i]` in the subset. Append `nums[i]` to `subset`, then recursively call `dfs(i + 1)`.
   - Exclude the current element `nums[i]` from the subset. Skip appending `nums[i]` to `subset` and recursively call `dfs(i + 1)`.
4. Initialize an empty subset and start the DFS traversal from index 0.
5. Return the resulting array of subsets.

## Time Complexity
The time complexity of this approach is O(2^N), where N is the number of elements in the input array. This is because each element has two choices: either to include it or exclude it from a subset, and there are 2^N possible subsets.

## Space Complexity
The space complexity is O(N * 2^N), where N is the number of elements in the input array. This is because there can be a maximum of 2^N subsets, each containing up to N elements.


## Remark

The line `subset = subset[:len(subset)-1]` is significant because it removes the last element from the subset. In the context of backtracking algorithms, this line is used to backtrack or undo the inclusion of the current element in the subset before exploring other possibilities.

Here's why this line is important:

1. **Backtracking**: After including an element in the subset and exploring all possible combinations with that element included, we need to backtrack to explore other combinations where that element is not included. Backtracking involves removing the last element added to the subset so that we can explore other branches of the recursion tree.

2. **Preparation for Next Iteration**: In the recursive function `backtrack`, after processing the current element and its subsets, we remove the last element from the `subset` before moving on to the next iteration. This prepares the `subset` for the next recursive call, ensuring that we correctly build subsets without the current element for the subsequent iterations.

3. **Maintaining State**: By removing the last element from the `subset`, we revert it to the state it was in before adding the current element. This ensures that we explore all possible combinations of elements without duplicates and without missing any combination.

In summary, the line `subset = subset[:len(subset)-1]` plays a crucial role in backtracking algorithms by allowing us to backtrack and explore different combinations of elements while constructing subsets.

### MISTAKE

Refer `pre-req.go` file to understand the mistake made

Suppose we have a slice `subset := []int{1, 2}`.

1. `res = append(res, append([]int{}, subset...))`: Here, we are creating a new slice with the elements of `subset` (1 and 2) using the variadic `...` syntax, and then appending that new slice to `res`. The inner `append([]int{}, subset...)` creates a new slice with the same elements as `subset`, ensuring that modifications to `subset` don't affect the appended slice in `res`. This way, `res` contains copies of the subsets, not references to the original `subset`.

2. `res1 = append(res1, subset)`: Here, we are directly appending the `subset` slice to `res1`. Since slices in Go are reference types, this means `res1` will hold a reference to the same underlying array as `subset`. Any modifications made to `subset` will reflect in `res1` because they both point to the same underlying array. This can lead to unexpected behavior if you modify `subset` after it has been appended to `res1`.




### Output

```
go run alt.go
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
