# Baseball Game

## Problem

You are keeping score for a baseball game with strange rules. The game consists of several rounds, where each round you get a score. The scores are represented by a list of strings, where each string represents one round of the game.

At the beginning of each round, you may get one of the following actions:

1. An integer `x` - Represents a round's score of `x` points.
2. "+" - Represents that the points you get in this round are the sum of the last two valid round's points.
3. "D" - Represents that the points you get in this round are twice the last valid round's points.
4. "C" - Represents the last valid round's points you get were invalid and should be removed.

Each round's operation is permanent and could have an impact on the round before and the round after.

You need to return the sum of the scores after all the rounds.

## Approach:
The function iterates through the given list of operations and maintains a record of valid round scores. It uses a switch statement to handle different types of operations:

- For integer scores, it converts the string to an integer and appends it to the record.
- For "+", it adds the sum of the last two valid round's scores to the record.
- For "D", it doubles the last valid round's score and appends it to the record.
- For "C", it removes the last valid round's score from the record.

Finally, it sums up all the valid round scores in the record and returns the result.

## Time Complexity:
O(n)

## Space Complexity:
O(n)

