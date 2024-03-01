# Last stone Weight

You are given an array of integers `stones`, where `stones[i]` is the weight of the `i`th stone. You are playing a game with these stones. On each turn, you choose the two heaviest stones and smash them together. If the weights of the two stones are equal, both stones are destroyed. Otherwise, the heavier stone is destroyed, and the lighter one's weight is reduced by the weight of the heavier stone. The game continues until there is at most one stone left. Return the weight of the last remaining stone.

## Intuition
To efficiently find and remove the two heaviest stones in each turn, we can use a max heap (priority queue). By repeatedly smashing the two heaviest stones until there is at most one stone left, we can determine the weight of the last remaining stone.

## Approach
1. Implement a max heap data structure to store the stone weights.
2. Push all the stone weights into the max heap.
3. Repeat the following steps until there is at most one stone left:
   - Pop the two heaviest stones from the max heap.
   - If the weights are equal, both stones are destroyed.
   - Otherwise, the heavier stone is destroyed, and the lighter one's weight is reduced by the weight of the heavier stone.
   - If a stone's weight becomes zero, it is removed from the game.
4. If there are no stones left, return 0. Otherwise, return the weight of the last remaining stone.

## Time Complexity
- Pushing all the stone weights into the max heap takes O(n log n), where n is the number of stones.
- Each turn involves popping two elements from the max heap, which takes O(log n) time.
- Since each stone may be removed from the heap at most once, the total time complexity is O(n log n), where n is the number of stones.

## Space Complexity
O(n) where n is the number of stones. This space is primarily used to store the stone weights in the max heap