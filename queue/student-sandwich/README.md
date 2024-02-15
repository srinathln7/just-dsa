# Count Students

## Problem

You are given two integer arrays `students` and `sandwiches`, where `students[i]` is the preference of the `ith` student and `sandwiches[j]` is the type of the `jth` sandwich. The types of sandwiches are represented by integers, where `0` represents circular sandwiches and `1` represents square sandwiches. You want to determine the maximum number of students that can be satisfied with their sandwiches.

## Approach

We can solve this problem by counting the number of students who prefer circular and square sandwiches. Then, we iterate through the sandwiches and check if there are enough sandwiches left for the students' preferences. We decrement the count of available sandwiches as students take sandwiches, and return the remaining count of unsatisfied students.

## Implementation

We use an array `count` of size 2 to store the count of students who prefer circular and square sandwiches. We iterate through the `students` array to count the preferences. Then, we iterate through the `sandwiches` array and decrement the count of available sandwiches as students take sandwiches. Finally, we return the sum of remaining unsatisfied students.

## Time Complexity

The time complexity of this approach is O(n), where n is the length of the `students` array or the `sandwiches` array.

## Space Complexity

The space complexity is O(1) since we are using a fixed-size array to store counts.
