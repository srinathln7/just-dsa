# Partition Label
Given a string `s`, partition it into as many parts as possible so that each letter appears in at most one part. Return a list of integers representing the length of these partitions.

## Intuition
The intuition behind this problem is to partition the string such that each letter appears in at most one part. To achieve this, we can keep track of the last occurrence index of each character in the string. Then, we iterate through the string and update the last occurrence index of the current character. Whenever we find a character whose last occurrence index equals the current index, we know that all characters in the substring from the last partition to the current index appear only in this partition.

## Approach
1. Create a map `lastIdxSet` to store the last occurrence index of each character in the string.
2. Iterate through the string and update the last occurrence index of each character in the `lastIdxSet`.
3. Initialize variables `left` and `lastIdx` to track the start index of the current partition and the last occurrence index of any character in the current partition, respectively.
4. Iterate through the string using a sliding window approach:
   - For each character, update `lastIdx` if the last occurrence index of the character is greater than the current `lastIdx`.
   - If `lastIdx` is equal to the current index, it means that all characters in the substring from `left` to `right` appear only in this partition. Append the length of this partition to the `result` array and update `left` to `right + 1`.
5. Return the `result` array containing the lengths of the partitions.

## Time Complexity
- Let `n` be the length of the input string `s`.
- The time complexity of the `partitionLabels` function is O(n) because we iterate through the string once to populate the `lastIdxSet` and once again to partition the string.

## Space Complexity
- The space complexity of the `partitionLabels` function is O(1) because the space used by the `lastIdxSet` map and the `result` array is proportional to the size of the alphabet, which is constant.