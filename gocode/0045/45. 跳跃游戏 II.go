package l0045

// 动态规划
func jump1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	var cnt int
	for n := 1; n < len(nums); n++ {
		cnt = n
		for i := 0; i < n; i++ {
			if nums[i] >= n-i && dp[i] < cnt {
				cnt = dp[i]
			}
		}
		dp[n] = cnt + 1
	}

	return dp[len(nums)-1]
}

// 贪心
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	MaxPath := len(nums) - 1
	maxPath := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i >= MaxPath {
			return dp[i] + 1
		}
		for j := maxPath + 1; j <= nums[i]+i && j < len(nums); j++ {
			dp[j] = dp[i] + 1
		}
		if maxPath < nums[i]+i {
			maxPath = nums[i] + i
		}
	}
	return dp[MaxPath]
}
