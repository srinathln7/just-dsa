# Simplify Path - Implement "cd" command
Given an absolute path for a file (Unix-style), simplify it. The absolute path is a string that contains multiple components separated by a slash '/' character, where each component represents a folder name or a file name. The simplified path should have the following properties:
- The path is always absolute, starting with the root '/'.
- Any consecutive slashes are treated as a single slash '/'.
- The path does not end with a trailing '/'.
- The path only contains the directories on the path from the root directory to the target file or directory (i.e., no '.' or '..' components).

## Intuition
To simplify the path, we can use a stack data structure. We split the input path by the delimiter '/', iterate through each component, and apply the following rules:
1. If the component is empty or '.', we ignore it.
2. If the component is '..', we pop the top element from the stack if the stack is not empty.
3. Otherwise, we push the component onto the stack.
Finally, we reconstruct the simplified path by joining the components in the stack with '/' delimiter.

## Approach
1. Split the input path string by the delimiter '/' to obtain individual components.
2. Iterate through the components:
    - If the component is empty or '.', skip it.
    - If the component is '..', pop the top element from the stack if the stack is not empty.
    - Otherwise, push the component onto the stack.
3. Construct the simplified path by joining the components in the stack with '/' delimiter.
4. If the stack is empty, return the root directory '/'; otherwise, return the simplified path.

## Time Complexity
- Let n be the length of the input path string.
- Splitting the path string takes O(n) time.
- Iterating through the components and operations within the loop take O(n) time in total.
- Constructing the simplified path takes O(n) time.
- Therefore, the overall time complexity is O(n).

## Space Complexity
- The space complexity is O(n) because we use additional space to store the stack and split components.
