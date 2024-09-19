package l0072

import "fmt"

func minDistance(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}
	dp[0][0] = 0

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(getMin(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}

	fmt.Println(dp)

	return dp[len(text1)][len(text2)]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
