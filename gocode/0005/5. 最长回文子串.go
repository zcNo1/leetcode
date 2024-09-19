package l0005

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]bool, len(s)+1)
		dp[i][0] = true
		dp[i][1] = true
	}
	dp[0] = make([]bool, len(s)+1)
	dp[0][0] = true
	maxStr := s[0:1]
	maxLen := 1

	for n := 1; n <= len(s); n++ {
		for i := 2; i <= n; i++ {
			dp[n][i] = dp[n-1][i-2] && (s[n-i] == s[n-1])
			if dp[n][i] && i > maxLen {
				maxLen = i
				maxStr = s[n-i : n]
			}
		}
	}

	fmt.Println(dp)

	return maxStr
}
