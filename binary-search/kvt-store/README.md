# Time-vased KV Store

Design a time-based key-value data structure that can store multiple values for the same key at different timestamps and retrieve the key's value at a certain timestamp.

## Intuition
To store multiple values for the same key at different timestamps and efficiently retrieve the value associated with a certain timestamp, we can use a hashmap to map each key to a list of elements (each containing a value and a timestamp). We can then perform binary search on the list of elements associated with a key to find the value associated with the largest timestamp_prev <= timestamp.

## Approach
1. Define a `TimeMap` struct with a `kvt` field to store key-value pairs along with timestamps.
2. Implement a `Constructor` function to initialize a new TimeMap object.
3. Implement a `Set` method to store the key with the value at the given timestamp. Append a new element (containing value and timestamp) to the list associated with the key.
4. Implement a `Get` method to retrieve the value associated with the largest timestamp_prev <= timestamp for a given key. Perform binary search on the list of elements associated with the key to find the appropriate value.
5. Handle edge cases such as the key not existing or the timestamp not matching any value.
6. Demonstrate the usage of the TimeMap class with example calls to `Set` and `Get`.

## Time Complexity
- The time complexity of the `Set` method is O(1) for appending to a list.
- The time complexity of the `Get` method is O(log n), where n is the number of elements associated with the key.


