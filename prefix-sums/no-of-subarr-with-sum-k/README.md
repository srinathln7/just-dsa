
# Sub-array with sum k

Given an array `nums` of integers and an integer `k`, the task is to find the total number of subarrays whose sum equals `k`.

## Intution

The approach involves utilizing prefix sums and a hashmap to efficiently find subarrays whose sum equals the target k.

## Approach

1. **Prefix Sum and Hashmap:** 
   - The code initializes a hashmap called `prefixSum` to keep track of the cumulative sums encountered so far and their frequencies.
   - The key idea is to utilize the cumulative sum concept and hashmap to efficiently find subarrays with a sum equal to `k`.

2. **Base Case:**
   - The code sets `prefixSum[0] = 1` as a base case. This represents the situation where the cumulative sum up to the current index equals `k`, indicating an empty subarray.

3. **Iterating Through the Array:**
   - The code iterates through the `nums` array and calculates the cumulative sum `cumSum` up to the current index.
   
4. **Checking for Subarrays with Sum `k`:**
   - At each iteration, the code checks if the difference between the current cumulative sum `cumSum` and `k` exists in the `prefixSum` map.
   - If the difference exists in the map, it means there is a subarray with sum equal to `k` ending at the current index. The code increments the count by the frequency associated with that difference.
   
5. **Updating the `prefixSum` Map:**
   - After checking for subarrays with sum `k`, the code updates the `prefixSum` map by incrementing the frequency of the current cumulative sum.
   
6. **Returning the Result:**
   - Finally, the code returns the total count of subarrays with sum equal to `k`.

## Time Complexity:
- The time complexity of the code is O(n), where n is the length of the `nums` array. This is because the code iterates through the array once.

## Space Complexity:
- The space complexity of the code is O(n), where n is the length of the `nums` array. This is due to the usage of the `prefixSum` hashmap to store cumulative sums and their frequencies.

