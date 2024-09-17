package l0064

func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))

	dp[0] = grid[0][0]
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			dp[j] = getMin(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[len(grid[0])-1]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
