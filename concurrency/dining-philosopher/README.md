# Dining Philosophers
The Dining Philosophers problem is a classic concurrency problem in computer science. It involves a scenario where a number of philosophers sit at a dining table with a bowl of spaghetti and a fork between each pair of adjacent philosophers. The philosophers alternate between thinking and eating. To eat, a philosopher must acquire the forks to their left and right. However, the forks are shared resources, and multiple philosophers cannot use the same fork simultaneously, leading to potential deadlocks if not properly managed.

## Intuition
The goal is to design a solution that allows the philosophers to dine without causing deadlocks, starvation, or resource contention. Each philosopher must be able to pick up and put down the forks, ensuring that they can eat without interfering with each other.

## Approach
1. Initialize a set of forks and philosophers.
2. Each philosopher repeatedly thinks, picks up the forks when available, eats, and puts down the forks.
3. To avoid deadlocks, ensure that the philosophers always pick up the forks in the same order and release them in the reverse order.
4. Use synchronization mechanisms, such as mutex locks, to coordinate access to shared resources (i.e., the forks).
5. Allow the philosophers to dine concurrently to maximize efficiency.

## Time Complexity
The time complexity of the solution depends on the number of iterations each philosopher performs (`TOTAL_ITERATIONS`) and the time it takes for each philosopher to think and eat. Assuming each action takes constant time, the overall time complexity is O(TOTAL_PHILOSOPHERS * TOTAL_ITERATIONS).

## Space Complexity
The space complexity is determined by the number of philosophers (`TOTAL_PHILOSOPHERS`) and the number of forks (`TOTAL_FORKS`). Both are constant values specified in the problem. Thus, the space complexity is O(1).