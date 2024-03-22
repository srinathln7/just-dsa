Y# Top K frequent elements

Given an integer array `nums` and an integer `k`, return the top k frequent elements in the array.

## Intuition
To find the top k frequent elements, we can use a min heap to efficiently maintain the k most frequent elements encountered so far. By iteratively adding elements to the min heap while maintaining its size at most k, we can ensure that the min heap contains the top k frequent elements at any given time.

## Approach
1. Create a frequency table `freq` to store the count of each element in the array.
2. Initialize an empty min heap `minHeap`.
3. Iterate through the elements of the frequency table:
   - For each element (key-value pair) in the frequency table, push it onto the min heap as an `Element` struct containing the element value and its frequency.
   - If the size of the min heap exceeds k after adding an element, pop the smallest element from the min heap to maintain its size at most k.
4. After processing all elements, the min heap will contain the top k frequent elements.
5. Create a result slice and populate it with the values of the top k elements from the min heap.
6. Return the result slice containing the top k frequent elements.

## Time Complexity
Let n be the number of elements in the input array `nums` and m be the number of unique elements in `nums`.
- Constructing the frequency table takes O(n) time.
- Building the min heap with k elements takes O(m log k) time.
- Constructing the result array takes O(k) time.

Overall, the time complexity of the algorithm is dominated by the heap operations, resulting in O(n log k) time complexity.

## Space Complexity
The space complexity of the algorithm is O(m + k), where m is the number of unique elements in `nums` (size of the frequency table) and k is the size of the min heap. 

