# Design LRU Cache 

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: `get` and `put`.

- `get(key)`: Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
- `put(key, value)`: Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

## Intuition

To efficiently support both `get` and `put` operations in constant time, we can use a combination of a hash map and a doubly linked list.

- A hash map allows for constant time lookups, enabling quick retrieval of values associated with keys.
- A doubly linked list preserves the order of elements based on their recent access, allowing us to efficiently identify the least recently used item for eviction.

## Approach

1. **Doubly Linked List Implementation**:
   - Define a `Node` struct to represent each node in the linked list. Each node contains a key-value pair, as well as pointers to its previous and next nodes.
   - Define an `LRUCache` struct to manage the cache. It includes the cache capacity, a hash map for key-value lookup, and pointers to sentinel nodes representing the head and tail of the linked list.
   
2. **Initialization**:
   - Initialize the `LRUCache` with the provided capacity. Create an empty hash map and two sentinel nodes for the head and tail of the doubly linked list.
   - Connect the head and tail sentinels to form an empty linked list.
   
3. **Get Operation**:
   - When performing a `get` operation, check if the key exists in the cache.
   - If the key exists, retrieve the corresponding value from the hash map and move the corresponding node to the front of the linked list (indicating it was recently used).
   - If the key does not exist, return -1.
   
4. **Put Operation**:
   - When performing a `put` operation, first check if the key already exists in the cache.
   - If the key exists, update the corresponding value and move the corresponding node to the front of the linked list.
   - If the key does not exist:
     - Create a new node with the given key and value.
     - Insert the new node just before the tail sentinel in the linked list.
     - Add the new key-value pair to the hash map.
     - If the cache capacity is exceeded, remove the least recently used node (at the end of the linked list) and delete its key from the hash map.

5. **Helper Functions**:
   - Define helper functions `Remove` and `Insert` to manipulate nodes within the linked list.
     - `Remove`: Removes a specified node from the linked list.
     - `Insert`: Inserts a specified node just before the tail sentinel in the linked list.

6. **Instantiate and Use**:
   - Instantiate an `LRUCache` object with the desired capacity.
   - Use the `get` and `put` methods to interact with the cache, as required.

## Time Complexity

- Both `get` and `put` operations run in O(1) time complexity since they involve constant-time operations such as hash map lookups and linked list manipulations.

## Space Complexity

- The space complexity of the `LRUCache` structure is O(capacity) for storing the hash map and the linked list nodes.