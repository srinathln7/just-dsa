#  IPO
A company plans to maximize its total capital before its initial public offering (IPO) by selecting a subset of projects to work on. Each project has a profit and requires a minimum capital to start. LeetCode can only start at most `k` projects before the IPO. The goal is to maximize the total capital after finishing at most `k` distinct projects.

## Intuition
To maximize the total capital, we need to select projects that provide the highest profit while ensuring that we have enough capital to start them. We can use a greedy approach where we maintain a max heap of projects sorted by their profits and a min heap of projects sorted by their capital requirements. We iteratively select the project with the highest profit from the max heap and ensure that we have enough capital to start it by checking the min heap. After selecting a project, we update our capital and continue the process until we have started `k` projects or there are no more projects left to select.

## Approach
1. Create a struct `Project` with `capital` and `profit` attributes to represent each project.
2. Implement a max heap data structure to store projects based on their profits.
3. Implement a function to push projects into the max heap and maintain its properties.
4. Implement a function to pop projects from the max heap and maintain its properties.
5. Iterate through the given projects and create project instances, then push them into both the max heap and min heap (sorted by capital).
6. Sort the projects based on their capital requirements.
7. Iterate through the projects while there are projects available and we can still start `k` projects:
   - Push projects from the min heap to the max heap as long as their capital requirements are within our current capital.
   - Pop the project with the highest profit from the max heap.
   - Update our capital with the profit from the selected project.
   - Decrement `k` to track the number of projects started.
8. Return the final capital after starting `k` projects.

## Time Complexity
- Sorting the projects based on their capital requires O(n log n) time.
- Pushing and popping projects from the max heap requires O(log n) time.
- Overall, the time complexity is O(n log n).

## Space Complexity
- The space complexity is O(n) to store the projects and the max heap.