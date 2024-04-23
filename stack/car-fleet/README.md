# Car Fleet
You are given an array `position` and an array `speed`, both of length `n`, where `position[i]` is the position of the ith car and `speed[i]` is the speed of the ith car (in meters/second). You are also given an integer `target` which is the position of a target car.

The goal is to find out how many car fleets will arrive at the target location.

Each car fleet is a group of cars that travel at the same speed and arrive at the same position. Cars in a fleet cannot overtake each other. If a car in a fleet catches up to another car, it will join the fleet behind the other car.

## Intuition
The key idea behind this problem is to simulate the movement of cars and determine how many fleets will arrive at the target location. Since cars can join a fleet if they catch up with a slower car in front of them, we need to track the time it takes for each car to reach the target location. We can start from the back and simulate the movement of cars to find out which cars form a fleet.

## Approach
1. Create a `car` struct to represent each car, containing its position and speed.
2. Create an array of `car` structs based on the given positions and speeds.
3. Sort the cars based on their positions.
4. Iterate through the cars from right to left:
   - Calculate the time it takes for each car to reach the target location.
   - Push the time to reach the target into a stack.
   - If the penultimate car reaches the target before the last car, they must have collided and become one car fleet. In this case, remove the last element from the stack.
5. Return the length of the stack, which represents the number of car fleets.

## Time Complexity
- Let `n` be the number of cars.
- Sorting the cars based on their positions takes O(n log n) time.
- Iterating through the cars and calculating the time to reach the target takes O(n) time.
- Thus, the overall time complexity is O(n log n).

## Space Complexity
- The space complexity is O(n) for storing the cars and the stack of arrival times.