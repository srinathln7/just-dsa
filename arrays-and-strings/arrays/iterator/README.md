# Interleaving Iterators in 2D (Nested) List 
Given a list of lists `lists`, where each inner list represents a sequence of integers, implement the `interLeave` function to return a new list of lists with the elements interleaved from each input list.

## Intuition
The key idea is to use iterators for each inner list to iteratively extract elements one by one and interleave them into the result list.

## Approach
1. Create an `Iterator` struct to manage iterating over a list of integers.
2. Implement the `next()` method to return the next element in the list and advance the iterator.
3. Implement the `hasNext()` method to check if there are more elements in the list.
4. Initialize an array of iterators, one for each inner list in `lists`.
5. Iterate through the iterators, extracting one element from each iterator in each iteration until all elements are exhausted.
6. Append the extracted elements to the result list.
7. Return the resulting list of lists.

## Time Complexity
- Constructing the iterator array: O(n), where n is the number of inner lists.
- Interleaving elements: O(m), where m is the total number of elements across all lists.

## Space Complexity
- O(n) for the iterator array.
- O(m) for the resulting list of lists, where m is the total number of elements across all lists.