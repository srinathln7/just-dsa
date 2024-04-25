# Task Scheduler
You are given an array of CPU tasks, each represented by letters A to Z, and a cooling time, n. Each cycle or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least n intervals due to cooling time. Return the minimum number of intervals required to complete all tasks.

## Intuition
The key idea is to calculate the frequency of each task and sort them in descending order to find the maximum frequency (maxFreq). Then, we calculate the total number of idle slots required to space out tasks with the maximum frequency. After that, we optimize the number of idle slots by subtracting it based on the minimum of (task_freq, maxFreq - 1). If the number of idle slots becomes negative, we reset it to zero since negative values have no meaning in this context. Finally, we calculate the sum of all tasks and idle slots.

## Approach
1. Build the frequency set for each task.
2. Sort the frequency set in decreasing order.
3. Calculate the maximum frequency (maxFreq).
4. Calculate the total number of idle slots required in the worst-case scenario: (maxFreq - 1) * n.
5. Optimize the number of idle slots by subtracting based on the minimum of (task_freq, maxFreq - 1) for each task frequency.
6. Reset idle slots to zero if they become negative.
7. Return the sum of all tasks and idle slots.

## Time Complexity
Sorting the frequency set takes O(m log m) time, where m is the number of unique tasks. The overall time complexity is dominated by sorting.

## Space Complexity
The space complexity is O(m), where m is the number of unique tasks, to store the frequency set.

## Remark

Sure, let's consider a simple example to illustrate this.

Suppose we have the following tasks: A, B, and C, with frequencies as follows:
- A: 4 times
- B: 3 times
- C: 2 times

And let's assume the cooling time `n` is 2.

1. **Worst-case scenario calculation:** In the worst-case scenario, we would need to space out the tasks with the maximum frequency (task A in this case) by `n` intervals. Since task A occurs 4 times, we have 3 intervals between each occurrence where we cannot execute task A again due to the cooling time. So, the total number of idle slots required for task A would be `(4 - 1) * 2 = 6`.

2. **Optimization:** Now, we need to consider the frequencies of the other tasks (B and C) and use any extra idle slots to minimize the overall idle time.

   - For task B: Since task B occurs 3 times, we want to ensure that we have enough idle slots to accommodate its occurrences. However, since task B occurs fewer times than task A, we can't utilize all the idle slots reserved for task A for task B as well. We take the minimum of the frequency of task B (3) and `maxFreq-1` (4-1=3), which is 3. So, we subtract 3 from the total idle slots.
   
   - For task C: Similarly, we consider task C. Since task C occurs 2 times, and we've already allocated some idle slots for tasks A and B, we take the minimum of the frequency of task C (2) and `maxFreq-1` (3), which is 2. So, we subtract 2 from the remaining idle slots.

After this optimization process, we may have some idle slots left, which are not needed for any task. We ensure that the number of idle slots is non-negative by setting it to zero if it becomes negative.

Finally, the total number of intervals required to execute all tasks with the given cooling time is the sum of the lengths of the tasks and the optimized idle slots.