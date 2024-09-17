package l0416

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum == 0 {
		return true
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 0; j < nums[i-1] && j <= target; j++ {
			dp[i][j] = dp[i][j] || dp[i-1][j]
		}
		for j := nums[i-1]; j <= target; j++ {
			dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
		}
	}

	return dp[len(nums)][target]
}
