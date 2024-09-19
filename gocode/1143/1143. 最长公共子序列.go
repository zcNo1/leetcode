package l1143

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = getMax(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	//fmt.Println(dp)

	return dp[len(text1)][len(text2)]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
