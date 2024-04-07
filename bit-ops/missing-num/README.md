# Missing number
Given an array `nums` containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

## Intuition
We can use the properties of XOR to find the missing number. XOR is commutative and associative, and XORing a number with itself gives zero. XORing a number with zero gives the number itself.

## Approach
1. Initialize a variable `missing` to 0.
2. Iterate through the array `nums`.
3. For each element `num` at index `i`, XOR `num` with `missing` and XOR it with `(i + 1)`.
4. The final value of `missing` will be the missing number.
5. Return the missing number.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the input array `nums`. We iterate through the array once.

## Space Complexity
The space complexity is O(1) as we use only a constant amount of extra space.
