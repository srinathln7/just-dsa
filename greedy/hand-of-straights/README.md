# Hand of Straight
Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size `groupSize`, and consists of `groupSize` consecutive cards.
Given an integer array `hand` where `hand[i]` is the value written on the ith card and an integer `groupSize`, return `true` if she can rearrange the cards, or `false` otherwise.

## Intuition
To determine if Alice can rearrange the cards into groups of size `groupSize`, we need to ensure that each group contains consecutive cards. We can achieve this by first sorting the cards and then checking if consecutive cards are available for each group.

## Approach
1. Check if the number of cards is divisible by `groupSize`. If not, return `false`.
2. Build a frequency map `freq` to count the occurrences of each card value in the `hand` slice.
3. Sort the `hand` slice in ascending order.
4. Iterate through each card value in the sorted `hand` slice.
5. For each card value, check if consecutive cards are available to form a group of size `groupSize`.
6. If any group cannot be formed, return `false`.
7. If all groups can be formed, return `true`.

## Time Complexity
The time complexity is O(n log n) due to sorting, where n is the number of cards in the hand. The subsequent iteration to form groups of size `groupSize` takes linear time.

## Space Complexity
The space complexity is O(n) to store the frequency of each card value in the `freq` map.
