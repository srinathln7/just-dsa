package main

func change(amount int, coins []int) int {

    // Key Idea: Bottom-up DP. Variaton of coin change 1 problem. Here dp[i] represents the number of ways
    // you can sum up to a target amount `i` using the given `coins`. The trick is we solve the problem 
    // by asking a simple question, For ex: amount = 5, coins = [1,2,5], we ask ourselves, how many ways
    // we can make amount from 0 to 5 using just coin `1`. Then we ask, how many ways, we can make amount
    // from 0 to 5 using coins [1,2]. Then we answer the bigger problem of how many ways, we can sum up
    // to amount 0 to 5 using coins [1,2,5]. And we return dp[5] which represents the number of ways, we
    // can sum up to 5 using the given coins [1,2,5]

    dp := make([]int, amount+1)

    // Base cae: No. of ways you can make an amount of 0 using the different coins is only one way
    // Just dont choose any coins
    dp[0] = 1
  
    for _, coin := range coins {
        for targetAmt :=1; targetAmt <= amount; targetAmt++ {

            // We can add up to the target amount only the current coin value is lesser than or equal to the target amout
            // For ex: You can NEVER make up a target value of 5 euro by using a 10 euro coin           
            if coin <= targetAmt {
                    dp[targetAmt] += dp[targetAmt-coin]    
                }    
            }
        }
    

    return dp[amount]
}
