# Design Circular Queue
Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

## Intuition
To implement a circular queue, we use a linked list structure where each node contains a value and pointers to the next and previous nodes. We also maintain pointers to the head and tail of the circular queue. The circular nature of the queue is achieved by connecting the tail to the head and the head to the tail. We track the length and capacity of the queue to determine if it is full or empty.

## Approach
1. **Constructor**: Initialize the circular queue with a given capacity `k`. Create sentinel nodes for the head and tail, and connect them to each other to form a circular structure.
2. **EnQueue**: Add a new node with the given value to the end of the circular queue. Update the tail pointer and increase the length of the queue.
3. **DeQueue**: Remove the node at the front of the circular queue. Update the head pointer and decrease the length of the queue.
4. **Front**: Return the value of the node at the front of the circular queue.
5. **Rear**: Return the value of the node at the end of the circular queue.
6. **IsEmpty**: Check if the circular queue is empty by comparing the length to zero.
7. **IsFull**: Check if the circular queue is full by comparing the length to the capacity.

## Time Complexity
- EnQueue: \(O(1)\)
- DeQueue: \(O(1)\)
- Front: \(O(1)\)
- Rear: \(O(1)\)
- IsEmpty: \(O(1)\)
- IsFull: \(O(1)\)

## Space Complexity
- \(O(n)\), where \(n\) is the capacity of the circular queue.