# Permutations with no duplicates

Given a collection of numbers `nums`, which might contain duplicates, return all possible unique permutations in any order.

## Intuition

To generate all possible unique permutations, we can utilize backtracking. By recursively constructing permutations, ensuring that duplicates are handled properly, we can generate all valid permutations.

## Approach

1. **Frequency Table**: Build a frequency table to count the occurrences of each number in the input array `nums`.
2. **Backtracking (DFS)**: Implement a depth-first search (DFS) with backtracking to generate permutations.
    - For each number in the frequency table, decrement its count and add it to the current permutation.
    - Recur with the updated permutation and frequency table.
    - Upon reaching a permutation of the same length as `nums`, add it to the result.
    - Backtrack by restoring the frequency count and removing the last added number from the permutation.
3. **Result**: Return the resulting list of unique permutations.

## Time Complexity

- Let \( n \) be the length of the input array `nums`.
- Constructing the frequency table takes \( O(n) \) time.
- Generating all unique permutations using backtracking has a time complexity of \( O(n \cdot n!) \) since there are \( n! \) permutations, each requiring \( O(n) \) time to construct.
- Therefore, the overall time complexity is \( O(n \cdot n!) \).

## Space Complexity

- The space complexity is \( O(n \cdot n!) \) to store all the permutations generated.
- Additionally, the frequency table requires \( O(n) \) space.
- Thus, the overall space complexity is \( O(n \cdot n!) + O(n) \), which simplifies to \( O(n \cdot n!) \).