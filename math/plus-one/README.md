# Plus One
You are given a large integer represented as an integer array `digits`, where each `digits[i]` is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's. Increment the large integer by one and return the resulting array of digits.

## Intuition
To increment a large integer represented as an array of digits, we start from the least significant digit (the last element of the array) and move towards the most significant digit (the first element of the array). We add 1 to the least significant digit and propagate the carry if necessary. If the carry propagates to the most significant digit, we need to append a new digit to the array.

## Approach
1. Start iterating from the last digit towards the first digit of the array.
2. Increment the current digit by 1.
3. If the incremented digit becomes 10, set it to 0 and propagate the carry by adding 1 to the next digit.
4. If the carry propagates to the most significant digit, append a new digit (1) to the beginning of the array.
5. Continue iterating until the carry is 0 or all digits are processed.
6. Reverse the array to restore the original order of digits.
7. Return the resulting array.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of digits in the input array `digits`.

## Space Complexity
The space complexity is O(1) if the resulting array is returned in-place, and O(n) if a new array is created to store the result.
