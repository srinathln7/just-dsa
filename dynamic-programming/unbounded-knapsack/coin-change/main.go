package main 

func coinChange(coins []int, amount int) int {
    
    // Key Idea: We can use a bottom-up dynamic programming approach where dp[i] represents the fewest number of coins
    // required to make up amount `i`. We initialize the `dp` to a random value greater than the target amount which
    // is never attainable. Note this is still a brute force method, where we evaulate all possible values using all
    // possible coins. We use memoization technique to avoid repeated work.

    dp := make([]int, amount+1)
    for i :=0; i < amount+1; i++ {
        dp[i] = amount+1
    }  

    // Base case : The fewest number of coins to make a target amount=0 is 0
    dp[0] = 0

    // Now, we loop over the each and every target amount from [0...amount] and find the fewest number of coins
    // required to make up that amount
    for targetAmt :=1; targetAmt <= amount; targetAmt++ {
        for _, coin := range coins {

            // We can add up to the target amount only the current coin value is lesser than or equal to the target amout
            // For ex: You can NEVER make up a target value of 5 euro by using a 10 euro coin
            if coin <= targetAmt {

                // Decision to exclude or include the current coin.  If we include the current coin 
                //in consideration, then we add that coin and look for its complimentary  
                dp[targetAmt] = min(dp[targetAmt], 1 + dp[targetAmt-coin])
            }
        }
    } 

    // Despite running a bottom-up DP, it the value remains unchanged, then it means
    // it is not possible to obtain the given amount with any of the coins
    if dp[amount] == amount+1 {
        return -1
    }

    return dp[amount]
}
