# Merge K-sorted Linked List
Merge k sorted linked lists and return it as one sorted list.

## Intuition

A straightforward approach to merging k sorted lists would be to merge two lists at a time repeatedly until we have merged all lists into one. However, a more efficient approach is to divide the k lists into halves recursively and merge them using ideas from merge sort. This approach results in a time complexity of \(O(n \log k)\), where \(n\) is the total number of nodes across all lists.

## Approach

1. Implement a function `mergeKLists(lists []*ListNode)` that takes an array of pointers to ListNode (representing the head nodes of k lists) and merges them into one sorted list.
2. Base cases:
    - If the array `lists` is empty, return nil.
    - If the length of `lists` is 1, return the only list.
    - If the length of `lists` is 2, merge the two lists using the `mergeTwoLists` function.
3. If the length of `lists` is greater than 2, recursively divide the array into two halves and merge them.
4. Implement a helper function `mergeTwoLists(list1, list2 *ListNode)` that merges two sorted lists into one sorted list.
5. Iterate over both lists simultaneously, comparing the values of the current nodes and appending the smaller node to the merged list.
6. Once one list is exhausted, append the remaining nodes of the other list to the merged list.
7. Return the merged list.

## Time Complexity

- \(O(n \log k)\), where \(n\) is the total number of nodes across all lists and \(k\) is the number of lists. The time complexity arises from recursively dividing the lists into halves and merging them using the `mergeTwoLists` function.

## Space Complexity

- \(O(\log k)\) - The space complexity arises from the recursion stack depth when recursively dividing the lists into halves. Additionally, the space complexity of the `mergeTwoLists` function is \(O(1)\) since we only use a constant amount of extra space to merge the lists.