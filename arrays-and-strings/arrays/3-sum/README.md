# [3Sum](https://leetcode.com/problems/3sum/)
Given an array of integers `nums`, return all the unique triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`. Notice that the solution set must not contain duplicate triplets.



## Intuition:
To find unique triplets that sum up to zero, we can leverage the two-pointer approach along with a hash map. By iterating through the array and treating each element as a potential target sum, we can transform the problem into a two-sum problem. Sorting the array beforehand ensures that we can efficiently avoid duplicate triplets.

## Approach:
1. Sort the input array `nums` to ensure that we avoid duplicate triplets.
2. Iterate through each element `num` in the sorted array.
3. For each `num`, find all unique pairs of numbers in the subarray `nums[i+1:]` that sum up to `-num`, effectively reducing the problem to a two-sum problem.
4. Store the unique triplets in a hash map to avoid duplicates.
5. Return the unique triplets as the result.

## Time Complexity:
- Sorting the array takes O(n * log n) time.
- Finding two sums for each element takes O(n^2) time.
- Overall, the time complexity is O(n^2) due to sorting and the nested loop for finding two sums.

## Space Complexity:
- The space complexity is O(n^2) due to the hash map storing unique triplets.

