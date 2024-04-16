# Same tree
Given the roots of two binary trees `p` and `q`, determine if they are identical, i.e., if they have the same structure and the same node values.

## Intuition
The intuition behind this problem is to recursively compare the nodes of both trees. For two trees to be identical, the following conditions must be met:
1. Their root values must be equal.
2. Their left subtrees must be identical.
3. Their right subtrees must be identical.

## Approach
1. Define a function `isSameTree` that takes pointers to the roots of two binary trees `p` and `q` as input and returns a boolean indicating whether they are identical.
2. In the `isSameTree` function:
   - Base cases:
     - If both `p` and `q` are nil, return true (two empty trees are identical).
     - If one of `p` or `q` is nil and the other is not, return false (trees with different structures are not identical).
   - Recursively compare the roots and subtrees of both trees:
     - Check if the values of the current nodes (`p.Val` and `q.Val`) are equal.
     - Recursively call `isSameTree` for the left subtrees of `p` and `q`.
     - Recursively call `isSameTree` for the right subtrees of `p` and `q`.
   - Return true only if all the above conditions are met.
3. Call the `isSameTree` function with the roots of the two binary trees to determine if they are identical.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the smaller of the two binary trees. This is because we need to traverse each node once to compare their values.

## Space Complexity
The space complexity of this approach is O(h), where h is the height of the binary trees. This space is used for the recursive call stack. In the worst case, when the trees are skewed (i.e., height is equal to the number of nodes), the space complexity can be O(n).