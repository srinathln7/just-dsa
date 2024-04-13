# Line Sweep Algorithm - Number of Flowers in Full Bloom
You are given a 0-indexed 2D integer array flowers, where flowers[i] = [starti, endi] means the ith flower will be in full bloom from starti to endi (inclusive). You are also given a 0-indexed integer array people of size n, where people[i] is the time that the ith person will arrive to see the flowers.

Return an integer array answer of size n, where answer[i] is the number of flowers that are in full bloom when the ith person arrives.

## Intuition
To solve this problem, we can use a min-heap to maintain the order of events based on their time. We iterate over each flower and each person, pushing their arrival and departure times onto the min-heap. As we process each event from the min-heap, we update the count of flowers in bloom. Finally, we return the count of flowers in bloom for each person's arrival time.

## Approach
1. Implement a min-heap to store events (Plane) consisting of time and event type (open, visit, close).
2. Push the start and end times of flowers as open and close events onto the min-heap.
3. Push the arrival times of people as visit events onto the min-heap.
4. Initialize a count variable to track the number of flowers in bloom.
5. Initialize a map countSet to store the count of flowers in bloom for each person's arrival time.
6. While the min-heap is not empty:
   - Pop the top event from the min-heap.
   - If it's an open event, increment the count of flowers in bloom.
   - If it's a close event, decrement the count of flowers in bloom.
   - If it's a visit event, update the count of flowers in bloom for the corresponding person's arrival time in countSet.
7. Return the count of flowers in bloom for each person's arrival time from countSet.

## Time Complexity
- Let m be the total number of flowers and n be the number of people.
- Building the min-heap takes O(m log m + n log m) time.
- Processing events and updating counts takes O(m + n log m) time.
- Therefore, the overall time complexity is O(m log m + n log m).

## Space Complexity
- The space complexity is O(m + n) for storing events and countSet.
