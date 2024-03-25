# Jump 2
Given an array `nums` where `nums[i]` denotes the maximum jump length from index `i`, find the minimum number of steps needed to reach the last index (`n-1`) starting from index 0.

## Intuition
To minimize the number of steps needed, a greedy approach is adopted. We employ a 1D BFS (Breadth-First Search) to output the shortest path length from index 0 to index `n-1`.

## Approach
1. Initialize variables `n` (length of `nums`), `l` (left boundary), and `r` (right boundary) to track the current range of indices that can be reached with the current number of jumps.
2. Initialize `count` to track the number of steps taken.
3. Run a loop until the right boundary `r` is less than `n-1` (destination index).
4. Inside the loop, initialize `maxJump` to 0 to track the farthest we can jump from the current range.
5. Iterate through the current range `[l, r]` using a loop.
6. Calculate the current jump length from index `i` (`currJump`) by adding `i` to the value at index `i` in `nums`.
7. Update `maxJump` to the maximum of its current value and `currJump`.
8. Update `l` to `r + 1` and `r` to `maxJump` to move to the next range of indices that can be reached.
9. Increment `count` to track the number of steps taken.
10. Return `count`, which represents the minimum number of steps needed to reach the last index.

## Time Complexity
The time complexity of this approach is O(n) since we only iterate through the array once.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space regardless of the size of the input array `nums`.


## Alt.go

**Although all testcases pass, it results in timeout or memory limit exception.**


### Intuition
This problem can be solved using a Breadth-First Search (BFS) approach. We traverse the indices in the array, considering each index as a node in the graph. We explore all possible jumps from each index and keep track of visited indices to avoid revisiting them. Once we reach the destination index, we return the minimum number of steps taken.

### Approach
1. Define a struct `KV` to represent a key-value pair containing the index and its corresponding value.
2. Implement a function `Jump` that takes an integer array `nums` as input and returns the minimum number of steps needed to reach the last index.
3. Initialize variables:
    - `start`: a `KV` struct representing the start index with its value from `nums[0]`.
    - `dst`: a `KV` struct representing the destination index with its value from `nums[n-1]`.
    - `minlen`: an integer to track the minimum number of steps taken.
    - `visited`: a map to track visited indices.
    - `queue`: a queue to perform BFS traversal.
4. Perform BFS traversal:
    - Enqueue `start` into the queue.
    - While the queue is not empty, perform the following steps:
        - Dequeue a node `curr` from the queue.
        - If `curr` has been visited, continue to the next iteration.
        - Mark `curr` as visited.
        - If `curr` is equal to the destination (`dst`), return `minlen`.
        - Explore all possible jumps from `curr` by iterating over indices from `curr.idx+1` to `curr.idx+curr.val`.
            - If the index is within bounds and has not been visited, enqueue it into the queue.
        - Increment `minlen`.
5. If the destination is not reachable, return `-1`.

## Time Complexity
The time complexity of this approach depends on the size of the queue and the number of edges in the graph. In the worst case, each index is visited once, and each edge is explored once, resulting in a time complexity of O(n).

## Space Complexity
The space complexity is determined by the size of the queue and the visited map. In the worst case, all indices are enqueued into the queue and stored in the visited map, resulting in a space complexity of O(n).