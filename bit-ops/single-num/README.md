# Single Num
Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one. You must implement a solution with a linear runtime complexity and use only constant extra space.

## Intuition
To find the single element that appears only once in the array while maintaining a linear runtime complexity and using constant extra space, we can utilize the bitwise XOR operation. The key properties of the XOR operation are:
- XORing any number with itself results in 0.
- XORing any number with 0 results in the number itself.
- XORing is commutative and associative, meaning the order of operands does not matter.

## Approach
We iterate through the array and XOR each element with the result. Since XORing any number with itself results in 0, and XORing any number with 0 results in the number itself, the XOR operation effectively cancels out all pairs of identical elements, leaving only the single element that appears once. The final result is returned as the answer.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of elements in the array. This is because we iterate through the array once.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space regardless of the size of the input array.