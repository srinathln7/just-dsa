package main 

func mincostTickets(days []int, costs []int) int {
    
    // Key Idea: We build dp[i] using bottom-up dynamic programming approach where `dp[i]` represents 
    // min. cost of travelling until days[i]th day. For example, consider days = [1,4,6,7,8,20]
    // dp[1] represents the min.cost of travelling until 1st day. dp[2] represents the min. cost
    // travelling on days [1,4] i.e. days 1 and 4. Hence, dp[20] represents the min. cost of travelling
    // on days [1,4,6,7,8,20]

    // First, capture the days you are interested in travelling
    travel := make(map[int]bool)
    for _, day := range days {
        travel[day] = true
    }

    // Build a dp-array
    n := len(days)
    maxDay := days[n-1]
    dp := make([]int, maxDay+1)

    for day :=1; day <= maxDay; day++ {

        // If the day is not in your travel itineary, there is no specific cost associated with travelling on that day
        // Hence, the total cost still remains on your previous day we chose to travel. For example days=[1,4] - we dont
        // travel on day 2 and 3, hence the cost associated with them is simply the cost of travelling on day 1.
        if !travel[day] {
            dp[day] = dp[day-1]
            continue 
        }

        // If we are travelling on this day, we calculate the minimum by using all possible passes
        
        // Using 1 day pass that cost of travelling on that day is the sum of the previous day cost plus the 1-day fare on that day
        dp[day] = dp[day-1] + costs[0]

        // Using 7-day pass
        dp[day] = min(dp[day], dp[max(0, day-7)]+ costs[1])

        // Using 30-day pass
        dp[day] = min(dp[day], dp[max(0, day-30)]+ costs[2])
    }

    // Return the cost of travelling until the max. day on my travel itineary
    return dp[maxDay]
}
