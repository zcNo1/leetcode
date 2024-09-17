package l0139

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for _, s2 := range wordDict {
		wordMap[s2] = struct{}{}
	}

	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if _, ok := wordMap[s[j:j+i]]; ok {
				dp[i][j] = true
				continue
			}
			for k := 1; k < i; k++ {
				dp[i][j] = dp[i][j] || (dp[k][j] && dp[i-k][j+k])
			}
		}
	}

	return dp[len(s)][0]
}
