# Problem: Partition List

## Intuition:
Given a linked list and a value `x`, we need to partition the list such that all nodes less than `x` come before nodes greater than or equal to `x`. The relative order of the nodes inside both partitions should remain the same.

## Approach:
To solve this problem, we can create two separate linked lists, one for nodes less than `x` and the other for nodes greater than or equal to `x`. We traverse the original linked list and insert each node into either the left list (for nodes less than `x`) or the right list (for nodes greater than or equal to `x`). After partitioning, we merge the left and right lists to form the final partitioned list.

## Time Complexity:
The time complexity of this approach is O(n), where n is the number of nodes in the input linked list. We traverse the list once to partition it and then merge the two lists, which takes linear time.

## Space Complexity:
The space complexity is also O(n), where n is the number of nodes in the input linked list. We create two new linked lists to store the partitioned nodes, and the merged list can have at most n nodes.
