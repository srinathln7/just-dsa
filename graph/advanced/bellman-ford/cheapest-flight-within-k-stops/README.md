# Cheapest Flight With atmost `k` stops

You are given an array of flights indicating flights between cities with their prices. You need to find the cheapest price from a source city to a destination city with at most `k` stops. If there's no such route, return -1.

## Intuition: 

Bellman-Ford algorithm is used here because it can handle graphs with negative weights and constraints like at most `k` stops. We cannot do a straight forward implementation of Dijkstra's algorithm here because of the `atmost k stops` constraint.

## Approach
  1. Initialize a `prices` slice representing the price to reach each city from the source city. Initialize it to positive infinity except for the source city, which is set to 0.
  2. Iterate `k+1` times, representing `k` stops.
  3. In each iteration, update the prices based on the flights.
  4. If there's no change in prices after `k+1` iterations, it means there's no cheaper route from the source to the destination.

## Time Complexity
 O(k * |flights|), where |flights| is the number of flights.

## Space Complexity

 O(n), where n is the number of cities.

This approach efficiently handles the problem constraints and returns the cheapest price from the source city to the destination city with at most `k` stops, or -1 if no such route exists.