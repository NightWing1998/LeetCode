// Link - https://leetcode.com/problems/perfect-squares/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func numSquares(n int) int {
	if n < 4 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 3
	for i := 4; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i] != 0 {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			} else {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

// Time:
//  Usage: 24ms
// 	Beats: 80.67%
// Memory:
//  Usage: 5.8MB
// 	Beats: 70.97%