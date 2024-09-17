package l0062

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	dp := make([]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}
