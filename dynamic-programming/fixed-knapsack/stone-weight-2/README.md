# Stone Weight 2 
You are given an array `stones` where `stones[i]` represents the weight of the ith stone. We are tasked with finding the minimum difference between the weights of two piles of stones after smashing all stones.

## Intuition
This problem can be reframed as a fixed knapsack problem, where the objective is to divide the stones into two piles with (almost) equal weights. We calculate the total sum of the weights of all stones and divide it by 2 to get the target weight for each pile. Then, we use dynamic programming with memoization to find the minimum difference between the sum of the weights of the two piles and the target weight.

## Approach
1. Calculate the total sum of the weights of all stones.
2. Compute the target weight by dividing the total sum by 2.
3. Initialize a memoization map `dp` to store results for each state `(idx, total)`, where `idx` represents the current index of the stone being considered, and `total` represents the total weight of the stones in the current pile.
4. Define a recursive function `dfs` that takes `idx` and `total` as arguments and returns the minimum difference between the sum of the weights of the two piles and the target weight.
5. In the `dfs` function:
   - Base case:
     - If the current total weight exceeds or equals the target weight, return the absolute difference between the total weight and `(sum - total)` (the weight of the other pile).
     - If `idx` reaches the end of the array `stones`, return the absolute difference between the total weight and `(sum - total)`.
   - Check if the result for the current state `(idx, total)` exists in the memoization map. If it does, return the stored value.
   - Make two decisions:
     - Skip the current stone and move to the next index (`dfs(idx+1, total)`).
     - Include the current stone in the current pile (`dfs(idx+1, total+stones[idx])`).
   - Update the memoization map with the minimum value between the two decisions.
   - Return the result for the current state.
6. Return the result of the `dfs` function with initial parameters `idx` and `total` set to 0.

## Time Complexity
The time complexity of this solution is O(n * target), where n is the number of stones and target is the target weight for each pile.

## Space Complexity
The space complexity is O(n * target) due to the memoization map `dp`.
