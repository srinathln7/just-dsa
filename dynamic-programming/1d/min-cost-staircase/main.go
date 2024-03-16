package main

func minCostClimbingStairs(cost []int) int {

     // Key Idea: Bottom up dynamic programming: To get to the nth step, you have to find the min. cost of getting to
    // (n-1) and (n-2)th step 

    n := len(cost)
    
    // Initialize variables to store the minimum costs for the previous two steps
    prev1, prev2 := cost[0], cost[1]

    // Iterate through the remaining steps and update the minimum costs
    for i := 2; i < n; i++ {
        // Calculate the minimum cost to reach the current step
        current := min(prev1, prev2) + cost[i]

        // Update the minimum costs for the previous two steps
        prev1, prev2 = prev2, current
    }

    // Return the minimum of reaching the last two steps
    return min(prev1, prev2)
}


