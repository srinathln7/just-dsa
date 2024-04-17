# Combination Sum II

Given a collection of candidate numbers `candidates` and a target number `target`, find all unique combinations in `candidates` where the candidate numbers sum to `target`. Each number in `candidates` may only be used once in the combination, and the solution set must not contain duplicate combinations.

## Intuition:
To find unique combinations that sum to the target, we can use Depth First Search (DFS) with backtracking. We'll sort the input array `candidates` to avoid duplicate combinations. During DFS traversal, we'll skip duplicates by comparing each element with its previous element within the current combination.

## Approach:
1. Sort the input array `candidates`.
2. Define a DFS function `dfsWithBT` that takes `candidates`, `target`, `startIdx`, `comb`, and `result` as parameters.
3. In the DFS function:
   - If the `target` becomes 0, append the current combination to the result.
   - If the `target` becomes negative, return.
   - Iterate over the `candidates` starting from the `startIdx`.
   - Skip duplicates by comparing the current element with its previous element within the current combination.
   - Include the current number in the combination, update the `target`, and recursively call DFS with the updated `target` and `startIdx`.
   - Backtrack by removing the last added element from the combination.
4. Start DFS with an initial `startIdx` of 0, an empty combination, and the `result`.
5. Return the result containing all unique combinations.

## Time Complexity:
- Sorting the array takes O(NlogN) time.
- Generating combinations using DFS takes O(2^N) time.
- Overall time complexity: O(NlogN + 2^N), where N is the length of the `candidates` array.

## Space Complexity:
- The space complexity for storing combinations and recursion stack during DFS is O(2^N).
- Sorting the array requires O(N) space.
- Overall space complexity: O(2^N).