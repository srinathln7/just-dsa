package main 

func rob(nums []int) int {
    
    // Key Idea: This is a slight variation of the house robbery problem 1 where to track the max money we can rob without 
    // setting the alarm. We adopt a BOTTOM-UP DP approach similar to house robbery 1. The main difference is we have to 
    // assess both the options and take the maximum of two options. 
    // Opt1: Stealing from house 0 in clockwise without alerting the police.
    // Opt2: Stealing from house n-1 in anti-clockwise without alerting the police.
    
    n := len(nums)
    switch {
        case n == 0: return 0
        case n ==1: return nums[0]
    }

    prevMax := nums[0]
    currMax := max(nums[0], nums[1])
    for i :=2; i < n-1; i++ {
        tmp := currMax
        currMax = max(currMax, prevMax+nums[i])
        prevMax = tmp
    }
    maxMoneyForward := currMax

    prevMax = nums[n-1]
    currMax = max(nums[n-1], nums[n-2])
    for i := n-3; i >0; i-- {
        tmp := currMax
        currMax = max(currMax, prevMax + nums[i])
        prevMax = tmp
    }

    return max(maxMoneyForward, currMax)
}
