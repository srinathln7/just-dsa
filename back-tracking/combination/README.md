# Combination
You are given two integers `n` and `k`, where `n` represents the range `[1, 2, ..., n]`, and `k` represents the number of integers to choose from the range `n`. Return all possible combinations of `k` numbers from the range `[1, 2, ..., n]`.

## Intuition
This problem can be solved using backtracking. We can perform depth-first search (DFS) to generate all combinations of `k` numbers from the given range `[1, 2, ..., n]`.

## Approach
1. Create a function `combine` that takes `n` and `k` as input and returns a 2D slice representing all combinations.
2. Initialize an empty slice `result` to store the combinations.
3. Create a slice `nums` containing integers from `1` to `n`.
4. Call a recursive function `dfsWithBT` with parameters:
   - `nums`: The slice of integers `[1, 2, ..., n]`.
   - `comb`: An empty slice to store the current combination.
   - `k`: The number of integers to choose.
   - `startIdx`: The starting index for selecting integers.
   - `result`: A pointer to the slice `result` to store the final combinations.
5. In the recursive function `dfsWithBT`:
   - If the length of the combination `comb` is equal to `k`, append the combination to the `result`.
   - Iterate through the range `[startIdx, len(nums))`:
     - Append the current integer from `nums` to the combination `comb`.
     - Recur with the updated combination, incrementing the `startIdx` by 1.
     - Backtrack by removing the last element from the combination.
6. Return the `result`.

## Time Complexity
The time complexity of this approach is O(C(n, k) * k), where C(n, k) is the number of combinations of choosing `k` elements from `n`. Each combination takes O(k) time to construct.

## Space Complexity
The space complexity is O(C(n, k) * k), as we need to store all possible combinations of length `k` from the range `[1, 2, ..., n]`.