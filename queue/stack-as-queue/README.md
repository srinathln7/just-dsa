# Stack Implementation Using Queue

## Problem

Implement a stack using a queue data structure.

## Approach:
This implementation utilizes a single queue to simulate a stack. The key idea is to maintain the order of elements in the queue such that the front element always represents the top of the stack.

### Push Operation:
- To push an element onto the stack, simply enqueue the element to the queue.

### Pop Operation:
- To pop an element from the stack, dequeue and enqueue all elements in the queue except the last one, which represents the element to be popped. Then, dequeue and return the last element.

### Top Operation:
- The top operation returns the element at the front of the queue, which represents the top of the stack.

### Empty Operation:
- The empty operation checks if the queue is empty.

## Time Complexity:
- Push Operation: O(1) amortized time complexity. Although enqueueing the element might require dequeuing and enqueuing all elements in the queue, on average, this operation takes O(1) time.
- Pop Operation: O(n) time complexity, where n is the number of elements in the stack. This is because dequeuing and enqueuing all elements except the last one takes O(n) time.
- Top Operation: O(1) time complexity, as accessing the front element of the queue is a constant-time operation.
- Empty Operation: O(1) time complexity, as it simply checks if the queue is empty.

## Space Complexity:
- O(n), where n is the number of elements in the stack. This space is required to store the elements in the queue.

