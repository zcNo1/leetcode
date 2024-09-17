package l0005

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s)+1)
		dp[i][0] = true
		dp[i][1] = true
	}
	dp[0][0] = true
	maxStr := s[0:1]
	maxLen := 1

	for i := 1; i <= len(s); i++ {
		for j := 2; j <= i; j++ {
			dp[i][j] = dp[i-1][j] || dp[i-1][j-2] && (s[i-j] == s[i-1])
			if dp[i][j] && j > maxLen {
				maxLen = j
				maxStr = s[i-j:i]
			}
		}
	}

	fmt.Println(dp)

	return maxStr
}

[
[true true false false false false]
[true true false false false false]
[true true false false false false]
[true true false true false false]
[true true false true false false]
[true true false true false false]]
