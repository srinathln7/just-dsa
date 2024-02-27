# Combination Sum
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order. The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

# Intuition
This problem can be solved using backtracking. We start by considering each candidate element one by one. For each candidate, we have two choices: either include it in the current combination or skip it. We continue this process recursively until we reach the target or exhaust all candidates.

# Approach
We define a recursive function `generateCombinationSum` to generate all possible combinations. Within this function, we iterate over the candidates starting from the `startIdx`. For each candidate, we add it to the current combination, decrement the target by the candidate value, and recursively call the function with the updated parameters. After processing a candidate, we backtrack by removing it from the current combination to explore other possibilities. We stop recursion when the target becomes zero (indicating a valid combination) or negative (indicating an invalid combination).

# Time Complexity
The time complexity of this approach is exponential, as we generate all possible combinations. In the worst case, where all candidates are chosen, the time complexity can be O(2^n), where n is the number of candidates.

# Space Complexity
The space complexity is also exponential, as we store all possible combinations in the result array. In the worst case, the space complexity can also be O(2^n), where n is the number of candidates.
