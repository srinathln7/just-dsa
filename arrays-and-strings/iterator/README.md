**Problem:**

Given a list of integer lists, we need to interleave the elements vertically, preserving the order of elements in each list. Empty lists should be ignored.

**Intuition:**

The problem requires interleaving the elements vertically. We can achieve this by using iterators for each list. We iterate through each iterator, fetching the next element until all iterators are exhausted. We then append the fetched elements to the result slice. By repeating this process, we can interleave the elements vertically.

**Approach:**

1. Create a custom iterator struct to iterate through each list.
2. Initialize an empty result slice to hold the interleaved elements.
3. Initialize iterators for each list and store them in a slice.
4. Iterate through each iterator and fetch the next element until all iterators are exhausted.
5. Append the fetched elements to the result slice.
6. Repeat steps 4-5 until all iterators are exhausted.
7. Return the result slice.

**Time Complexity:**

Let n be the total number of elements across all lists.

- Initializing iterators: O(m), where m is the number of lists.
- Interleaving elements: O(n).

The overall time complexity is O(n).

**Space Complexity:**

- Space required for iterators: O(m), where m is the number of lists.
- Space required for result slice: O(n), where n is the total number of elements across all lists.

The overall space complexity is O(n).