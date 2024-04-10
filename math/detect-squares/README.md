# Detect Squares
Design an algorithm for a data structure called `DetectSquares`. This structure should allow adding points and counting the number of ways to choose three points from the data structure such that the three points and a given query point form an axis-aligned square with positive area.

## Intuition
We can store the points in a map where the key is a tuple of x and y coordinates and the value represents the count of how many times that point has been added. When adding a point, we update both the map and a slice containing all the points. When counting, we iterate over all points, check if they form a valid square with the query point, and calculate the count based on the occurrence of relevant points.

## Approach
1. Initialize a struct `DetectSquares` with two fields: `ptsMap` to store the count of points and `pts` to store all points.
2. Implement a constructor function `Constructor` to initialize the `DetectSquares` struct.
3. Implement the `Add` method to add a point to the data structure. Update both the `ptsMap` and `pts`.
4. Implement the `Count` method to count the number of valid squares with a given query point.
   - Iterate over all points in `pts`.
   - Check if the current point forms a valid square with the query point.
   - If yes, calculate the count based on the occurrence of relevant points in `ptsMap`.
5. Implement a helper function `abs` to calculate the absolute value of an integer.
6. Instantiate the `DetectSquares` object and call the `Add` and `Count` methods as required.

## Time Complexity
- Adding a point: O(1) for updating the map and slice.
- Counting: O(N) where N is the number of points in the structure, as we iterate over all points.

## Space Complexity
- O(N) where N is the number of points stored in the `pts` slice and the size of the `ptsMap`.