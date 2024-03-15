# Min. Cost for Travel Tickets

You are given an integer array `days` representing days of the year on which you plan to travel. Each day is an integer from 1 to 365. Train tickets are sold in three different ways:

1. A 1-day pass is sold for `costs[0]` dollars.
2. A 7-day pass is sold for `costs[1]` dollars.
3. A 30-day pass is sold for `costs[2]` dollars.

The passes allow that many days of consecutive travel. Return the minimum number of dollars needed to travel every day in the given list of days.

## Intuition
To minimize the cost of traveling, we need to compute the minimum cost for traveling up to each day. We can use dynamic programming to solve this problem. We will create an array `dp`, where `dp[i]` represents the minimum cost of traveling up to the `i`-th day. Then, we iterate through each day, considering the cost of using 1-day, 7-day, and 30-day passes, and update `dp[i]` accordingly.

## Approach
1. Initialize a map `travel` to keep track of the days you plan to travel. Set `true` for each day in the `days` array.
2. Create an array `dp` of size `maxDay + 1`, where `maxDay` is the maximum day in the `days` array.
3. Iterate through each day from 1 to `maxDay`.
    - If the day is not in the `travel` itinerary, set the cost for that day equal to the cost of the previous day.
    - If you are traveling on the current day, calculate the minimum cost by considering the cost of using 1-day, 7-day, and 30-day passes.
4. Return the minimum cost of traveling until the last day in the `days` array.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of days in the `days` array.

## Space Complexity
The space complexity is O(n), where n is the number of days in the `days` array, as we use an additional map and array to store the intermediate results.
