# Kth largest element in a stream

Design a data structure that efficiently maintains the k largest elements from a stream of integers.

## Intuition:
Maintaining the k largest elements from a stream requires a data structure that can efficiently handle insertions and removals while keeping track of the k largest elements at any given time.

## Approach:
1. **MinHeap Implementation:**
   - Utilize a min-heap to maintain the k largest elements.
   - Whenever a new element is added, insert it into the min-heap.
   - If the size of the heap exceeds k, remove the smallest element (root of the min-heap).
   - The root of the min-heap will always represent the kth largest element.
  
2. **Constructor Function:**
   - Initialize the KthLargest object with the given value of k and a slice of integers.
   - Iterate through the provided integers and add them to the min-heap.

3. **Add Method:**
   - Add a new integer to the KthLargest object.
   - Push the integer into the min-heap.
   - If the size of the min-heap exceeds k, remove the smallest element (maintaining k largest elements).
   - Return the root of the min-heap, which represents the kth largest element.

## Time Complexity:
- Construction: O(n log k), where n is the number of elements in the initial array and k is the given value.
- Add Operation: Push: O (log k) + Pop: O((n-k) log k) ~ O(m log(k)) where `m= n-k` and when `k` is small it will be O(nlog(k)) for inserting and adjusting the min-heap.

## Space Complexity:
- O(k) for maintaining the min-heap of size k.

This approach ensures that the k largest elements are efficiently maintained, allowing for constant-time access to the kth largest element.