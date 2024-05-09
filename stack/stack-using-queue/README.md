# Implement Stack using Queue
Implement a stack using queues. The stack should support the following operations: `Push`, `Pop`, `Top`, and `Empty`.

## Intuition
To implement a stack using queues, we can use two queues. We'll use one queue for the main storage of elements and another queue for auxiliary operations such as pushing elements onto the stack.

## Approach
1. Create a structure `MyStack` with a field `Queue` to store elements.
2. Implement the `Push` operation:
   - Append the new element to the main queue.
3. Implement the `Pop` operation:
   - Dequeue and enqueue elements from the main queue to the auxiliary queue until there's only one element left in the main queue.
   - Remove and return the last element from the main queue.
   - Swap the names of the main and auxiliary queues.
4. Implement the `Top` operation:
   - Return the last element of the main queue.
5. Implement the `Empty` operation:
   - Return true if the main queue is empty, false otherwise.

## Time Complexity
- Push operation: O(1)
- Pop operation: O(N), where N is the number of elements in the stack. Dequeueing and enqueuing elements to the auxiliary queue takes O(N) time.
- Top operation: O(1)
- Empty operation: O(1)

## Space Complexity
- O(N), where N is the number of elements in the stack, as we use extra space to store elements in the main queue.