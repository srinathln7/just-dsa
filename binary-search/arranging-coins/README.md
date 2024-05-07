# Arranging Coins

You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete. Given the integer n, return the number of complete rows of the staircase you will build.

## Intuition

The problem can be approached using binary search. We need to find the maximum value of k such that the sum of coins from 1 to k (inclusive) does not exceed n.

## Approach

1. Initialize two pointers, l and r, where l points to 1 and r points to n.
2. Use binary search to find the maximum value of k such that the sum of coins from 1 to k (inclusive) does not exceed n.
3. While l is less than or equal to r:
   - Calculate mid as l + (r-l)/2.
   - If n equals the sum of coins till k (sumTillK(mid)), return mid.
   - If n is less than the sum of coins till k, update r to mid - 1.
   - If n is greater than the sum of coins till k, update l to mid + 1.
4. Return r.

## Time Complexity

The time complexity of this approach is O(log n), as we are using binary search.

## Space Complexity

The space complexity is O(1) as we are using only a constant amount of extra space.

## Remark

In this binary search implementation, we are trying to find the maximum value of \( k \) such that the sum of coins from 1 to \( k \) (inclusive) does not exceed \( n \). 

- The variable \( l \) represents the lower bound of the range of possible values for \( k \), initialized to 1.
- The variable \( r \) represents the upper bound of the range, initialized to \( n \).
- At each step of the binary search, we update either \( l \) or \( r \) based on the comparison between \( n \) and the sum of coins till \( k \).
- When \( l \) becomes greater than \( r \), it indicates that we have found the maximum value of \( k \) where the sum of coins is less than or equal to \( n \).
- Therefore, we return \( r \) because it represents the maximum value of \( k \) that satisfies the condition.

Returning \( l \) instead of \( r \) would not give the correct result because \( l \) would be pointing to the first value where the sum exceeds \( n \), which is not the maximum valid value of \( k \).