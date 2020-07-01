func uniquePaths(m int, n int) int {
	if n == 0 || m == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][m-1] = 1
	}
	for j := 0; j < m; j++ {
		dp[n-1][j] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			dp[i][j] = dp[i][j+1] + dp[i+1][j]
		}
	}
	return dp[0][0]
}