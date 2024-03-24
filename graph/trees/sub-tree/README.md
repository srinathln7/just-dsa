# Sub-tree
Given the roots of a binary tree `root` and another binary tree `subRoot`, determine if `subRoot` is a subtree of `root`. A subtree of a binary tree `tree` is a tree that consists of a node in `tree` and all of this node's descendants. The tree `tree` could also be considered as a subtree of itself.

## Intuition
To determine if `subRoot` is a subtree of `root`, we can recursively compare each subtree of `root` with `subRoot`. If any subtree of `root` is identical to `subRoot`, then `subRoot` is considered a subtree of `root`.

## Approach
1. Define a function `isSubtree` that takes pointers to the roots of two binary trees `root` and `subRoot` as input and returns a boolean indicating whether `subRoot` is a subtree of `root`.
2. In the `isSubtree` function:
   - Check if `subRoot` is identical to `root` using the `isSameTree` function:
     - If `isSameTree(root, subRoot)` returns true, return true (since `subRoot` is identical to `root`).
   - If `subRoot` is not identical to `root`, recursively check if `subRoot` is a subtree of the left subtree of `root` or the right subtree of `root`:
     - If `root.Left` is not nil and `isSubtree(root.Left, subRoot)` returns true, return true.
     - If `root.Right` is not nil and `isSubtree(root.Right, subRoot)` returns true, return true.
   - If none of the above conditions are met, return false.
3. Define a helper function `isSameTree` that takes pointers to the roots of two binary trees `p` and `q` as input and returns a boolean indicating whether they are identical (same structure and node values).
4. In the `isSameTree` function:
   - Base cases:
     - If both `p` and `q` are nil, return true (two empty trees are identical).
     - If one of `p` or `q` is nil and the other is not, return false (trees with different structures are not identical).
   - Recursively compare the roots and subtrees of both trees:
     - Check if the values of the current nodes (`p.Val` and `q.Val`) are equal.
     - Recursively call `isSameTree` for the left subtrees of `p` and `q`.
     - Recursively call `isSameTree` for the right subtrees of `p` and `q`.
   - Return true only if all the above conditions are met.
5. Call the `isSubtree` function with the roots of the two binary trees to determine if `subRoot` is a subtree of `root`.

## Time Complexity
The time complexity of this approach depends on the size of the binary tree `root` and the binary tree `subRoot`. In the worst case, where `subRoot` is not a subtree of `root`, the time complexity can be O(m * n), where m is the number of nodes in `root` and n is the number of nodes in `subRoot`.

## Space Complexity
The space complexity of this approach depends on the depth of recursion, which is proportional to the height of the binary tree `root`. In the worst case, the space complexity can be O(h), where h is the height of the binary tree `root`.