# Asteroids Collision

We are given an array of integers `asteroids`, where each integer represents the size of an asteroid in a row. The absolute value of the integer represents its size, and the sign represents its direction: positive for right, negative for left. Each asteroid moves at the same speed.

When two asteroids meet, the smaller one will explode. If they are the same size, both will explode. Two asteroids moving in the same direction will never meet.

We need to find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

## Intuition

We can use a stack data structure to simulate the collisions between asteroids. As we iterate through the asteroids, we compare the current asteroid with the top of the stack to determine if a collision occurs. If a collision occurs, we handle it according to the given rules. If no collision occurs, we add the asteroid to the stack. After iterating through all the asteroids, the stack will contain the remaining asteroids after all collisions.

## Approach

1. Initialize an empty stack to store the remaining asteroids.
2. Iterate through each asteroid in the given array.
3. For each asteroid:
   - Compare it with the top of the stack to check for collisions.
   - If a collision occurs:
     - If the incoming asteroid has greater absolute size, remove the top asteroid from the stack.
     - If the incoming asteroid has smaller absolute size, discard the incoming asteroid.
     - If both asteroids have the same absolute size, remove both from the stack.
   - If no collision occurs, add the asteroid to the stack.
4. After iterating through all asteroids, the stack will contain the remaining asteroids after all collisions.

## Time Complexity

The time complexity of this approach is O(n), where n is the number of asteroids in the input array. We iterate through each asteroid once and perform constant-time operations within each iteration.

## Space Complexity

The space complexity is also O(n). In the worst case, when there are no collisions, all asteroids will be stored in the stack.

