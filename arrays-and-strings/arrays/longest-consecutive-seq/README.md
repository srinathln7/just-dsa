# [Longest consecutive sequence](https://leetcode.com/problems/longest-consecutive-sequence/description/):
Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

## Intuition:
The key idea is to identify the start of a consecutive sequence and then expand it to the right until no more consecutive elements are found. While iterating through the array, for each element, we check if its left neighbor exists. If not, we start counting the consecutive elements to the right.

## Approach:
1. Create a hashmap (`seqMap`) to store the presence of each number in the array.
2. Iterate through the array to populate the `seqMap`.
3. Initialize a variable `maxLen` to store the maximum length of consecutive elements sequence found so far.
4. Iterate through the array again. For each element:
   - Check if its left neighbor exists in `seqMap`. If not:
     - Initialize a counter `seqLen` to 1.
     - Increment the current number and check if the next consecutive element exists in `seqMap`.
     - If yes, increment `seqLen` and continue until no more consecutive elements are found.
     - Update `maxLen` with the maximum of its current value and `seqLen`.
5. Return `maxLen`.

## Time Complexity:
- Populating `seqMap`: O(n), where n is the number of elements in the input array.
- Iterating through the array twice: O(2n) ≈ O(n).
- Overall, the time complexity is O(n).

## Space Complexity:
- Space required to store `seqMap`: O(n), where n is the number of elements in the input array.
- Additional space used for variables: O(1).
- Overall, the space complexity is O(n).
