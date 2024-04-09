# Multiply Strings
Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string. Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

## Intuition
To multiply two numbers represented as strings, we can perform long multiplication just like we do manually. We multiply each digit of one number with each digit of the other number and accumulate the results in the appropriate positions. After all the digits are multiplied and accumulated, we perform carries to obtain the final result.

## Approach
1. Check if either `num1` or `num2` is "0". If so, return "0".
2. Reverse the strings `num1` and `num2`.
3. Initialize a result array to store the product, with length equal to the sum of lengths of `num1` and `num2`.
4. Iterate through each digit of `num1` and `num2` from right to left.
5. For each pair of digits, multiply them and accumulate the result in the appropriate position of the result array.
6. After multiplication, perform carries to adjust the result array.
7. Reverse the result array and remove leading zeros.
8. Convert the result array back to a string and return it.

## Time Complexity
The time complexity of this approach is O(m * n), where m and n are the lengths of the input strings `num1` and `num2`, respectively. This is because we iterate through each digit of `num1` and `num2` and perform multiplication.

## Space Complexity
The space complexity is O(m + n), where m and n are the lengths of the input strings `num1` and `num2`, respectively. This is because we store the resulting product in an array of length m + n.


## Remarks - String Reverse

The provided code attempts to reverse a string by swapping characters from the beginning and end of the string iteratively until the midpoint is reached. However, in Go, strings are immutable, meaning their contents cannot be modified after creation. Therefore, attempting to assign a value to an index of a string (`s[i] = s[n-1-i]`) will result in a compilation error. To fix this issue, you can convert the string to a slice of runes, which are mutable, perform the swapping, and then convert it back to a string. 

