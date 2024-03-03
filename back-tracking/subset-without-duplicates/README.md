# Subsets with no duplicates

Given an integer array `nums` that may contain duplicates, the task is to return all possible subsets (the power set) of `nums`. The solution set must not contain duplicate subsets.

**Intuition:**
To generate all possible subsets without duplicates, we can use Depth-First Search (DFS) with backtracking. We'll track the frequency of each element using a frequency map. During each iteration of DFS, we'll consider whether to include the current element in the subset based on its frequency and the elements already present in the subset.

**Approach:**
1. Initialize an empty result slice to store all subsets.
2. Create a frequency map (`freq`) to track the count of each element in `nums`.
3. Perform DFS with backtracking:
   - At each step, append the current subset to the result.
   - Iterate over each unique element (`num`) in the frequency map:
     - Skip forming the subset if either the frequency of `num` is 0 or the last element in the current subset is greater than `num`. This avoids duplicate subsets.
     - Otherwise, append `num` to the subset, decrement its count in the frequency map, and recursively call DFS.
     - After the recursive call, backtrack by removing `num` from the subset and restoring its count in the frequency map.
4. Return the result containing all generated subsets.

**Time Complexity:**
The time complexity of the DFS algorithm is O(2^N), where N is the number of elements in `nums`. This is because each element has two choices: either include it or exclude it from the subset.

**Space Complexity:**
The space complexity is also O(2^N) due to the recursive stack space and the space required to store the result.

