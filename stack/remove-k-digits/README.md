# Remove k digits to minimize the number
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

### Intuition
The idea is to use a monotonic stack technique to maintain the stack in increasing order. If the stack is in increasing order, we can remove digits from the right. If the numbers are in decreasing order, we remove digits from the left. We aim to maintain a monotonic stack of increasing order so that we can start popping elements as long as k is zero.

### Approach
1. Base case: If the string length is less than or equal to k, return "0" since we can remove all the characters.
2. Initialize an empty stack.
3. Iterate through each character in the string num.
4. For each character:
   - For a non-empty stack and if the incoming character is less than the stack's top, pop elements from the stack until the condition is satisfied or k becomes zero.
   - Append the character to the stack if it's not a leading zero.
5. After the iteration, if k is still non-zero and the stack is not empty, remove the remaining digits from the most significant bit (MSB) until either the stack is empty or k becomes zero.
6. Trim leading zeros from the result.
7. Return the result. If the result is empty, return "0".

### Time Complexity
The time complexity is O(n), where n is the length of the input string num.

### Space Complexity
The space complexity is O(n) due to the stack used to store the characters.