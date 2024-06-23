package l0064

import "log"

func minPathSum(grid [][]int) int {
	return getMinSum(grid, 0, 0, 0)
}

func getMinSum(grid [][]int, idx, idy, sum int) int {
	sum += grid[idx][idy]
	sumx := sum
	sumy := sum
	if idx+1 < len(grid) {
		sumx = getMinSum(grid, idx+1, idy, sum)
	}
	if idy+1 < len(grid[0]) {
		sumy = getMinSum(grid, idx, idy+1, sum)
	}
	if sumx == sumy {
		return sumx
	}
	if sumx == sum {
		return sumy
	}
	if sumy == sum {
		return sumx
	}
	if sumx < sumy {
		return sumx
	}
	return sumy
}

func minPathSum2(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] += dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] += dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = getMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	log.Println(dp)
	return dp[len(grid)-1][len(grid[0])-1]
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
