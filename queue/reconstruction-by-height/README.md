# Queue reconstruction based on height and number of people infront of them

Given an array of people, where each person is represented by their height and the number of people taller than or equal to them standing in front of them, reconstruct the queue based on this information.

**Intuition:**
To reconstruct the queue, we need to sort the people in such a way that taller people with fewer or equal number of people in front of them come first. We can achieve this by sorting the people array based on their heights in descending order and, for people with the same height, sorting them based on the number of people in front of them in ascending order.

**Approach:**
1. Sort the `people` array based on height in descending order.
2. If two people have the same height, sort them based on the number of people in front of them in ascending order.
3. Iterate over the sorted array and insert each person into the reconstructed queue at the appropriate position.
4. Use the `copy` function to shift elements to the right before inserting each person to maintain the correct order.

**Time Complexity:**
- Sorting the array takes O(n log n) time, where n is the number of people.
- Inserting each person into the queue takes O(n^2) time due to shifting elements using the `copy` function.
- Overall time complexity: O(n^2).

**Space Complexity:**
- The space complexity is O(n) for the reconstructed queue and O(1) for other variables.
- Overall space complexity: O(n).