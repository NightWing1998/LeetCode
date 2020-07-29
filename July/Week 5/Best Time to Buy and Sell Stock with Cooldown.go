// Link - https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func max(a,b int) int {
    if a > b {return a;}
    return b;
}

func maxProfit(prices []int) int {
    if(len(prices) <= 1) {
        return 0;
    }
    if len(prices) == 2{
        if prices[1] > prices[0] {return prices[1]-prices[0];}
        return 0;
    }
    
    dp := make([][2]int,len(prices));
    
    // i[0] --> sell margin
    dp[0][0] = 0;
    // i[1] --> buy margin
    dp[0][1] = -prices[0];
    
    dp[1][0] = max(dp[0][1] + prices[1],dp[0][0]);
    dp[1][1] = max(-prices[1],dp[0][1]);
    
    for i := 2; i < len(prices); i++ {
        // sell val is max of (bought yesterday sold today) and (sold yesterday)
        dp[i][0] = max(dp[i-1][1] + prices[i],dp[i-1][0]);
        // buy val is max of (sold before cooldown yesterday bought today) and (bought yesterday)
        dp[i][1] = max(dp[i-2][0] - prices[i],dp[i-1][1]);
    }
    
    return dp[len(prices)-1][0];
}

// Time - 0ms(beats 100%)
// Memory - 2.4MB(-)

