# First Unique Character in a String

## Problem Statement
Given a string `s`, find the first non-repeating character in it and return its index. If it does not exist, return -1.

## Intuition
This solution utilizes an array to count the occurrences of each character in the string. By traversing the string twice, we can determine the index of the first unique character. In the first pass, we count the occurrences of each character. In the second pass, we check each character's count and return the index of the first character with a count of 1.

## Approach
1. Initialize an array `count` of size 26 to store the count of each lowercase English alphabet character.
2. Traverse the string `s` and increment the count of each character encountered.
3. Traverse the string `s` again. For each character encountered, check its count in the `count` array. If the count is 1, return the index of the character.
4. If no unique character is found, return -1.

## Time Complexity
The solution requires traversing the string twice, resulting in a time complexity of O(n), where n is the length of the string.

## Space Complexity
The solution utilizes an array of size 26 to store the count of characters, resulting in a space complexity of O(1), as the size of the array is constant.
