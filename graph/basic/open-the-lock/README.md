# Open the Lock
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around. Each move consists of turning one wheel one slot. The lock initially starts at '0000', representing the state of the 4 wheels.

You are given a list of deadends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it. Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

## Approach
The problem is equivalent to finding the shortest path from the initial state ("0000") to the target state while avoiding deadends. We can use a breadth-first search (BFS) approach to efficiently solve this problem.

1. **Initialize Data Structures**:
   - Create a set `deadendSet` to store the deadends for efficient lookup.
   - Check if the initial state is a deadend. If so, return -1.
   - Initialize a map `visited` to keep track of visited states during BFS traversal.
   - Initialize a count variable to keep track of the number of turns.

2. **Breadth-First Search (BFS)**:
   - Start BFS traversal from the initial state ("0000").
   - Enqueue the initial state into a queue.
   - While the queue is not empty:
     - Process all nodes at the current level:
       - Dequeue the current node.
       - Skip if the current node is already visited.
       - Mark the current node as visited.
       - Check if the current node is the target state. If so, return the count.
       - Generate child nodes by turning each wheel clockwise and anti-clockwise.
       - Enqueue child nodes into the queue if they are not visited and not in deadends.
     - Move to the next level of BFS traversal.

3. **Helper Function to Turn Wheels**:
   - Define a helper function `turnWheels` to simulate turning a wheel at a given slot in a given direction (clockwise or anti-clockwise).
   - Handle wrapping around when turning the wheels.

## Time Complexity
The time complexity of this approach is O(N), where N is the total number of possible states reachable from the initial state. Since each wheel has 10 slots, there are 10^4 possible states (10 choices for each of the 4 wheels), leading to a manageable search space.

## Space Complexity
The space complexity of this approach is O(N), where N is the maximum size of the queue during BFS traversal. In the worst case, all possible states reachable from the initial state need to be stored in the queue. However, the actual space usage may be lower depending on the structure of the lock and the number of deadends.