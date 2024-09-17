package l0055

// 动态规划
// dp[n] = for 0<=i<n || (dp[i]&&(nums[i]>=n-i))
// dp[0] = true // 第一个位置必定可以到达，为true
func canJump1(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true

	for n := 1; n < len(nums); n++ {
		for i := 0; i < n; i++ {
			dp[n] = dp[n] || (dp[i] && nums[i] >= n-i)
		}
	}

	return dp[len(nums)-1]
}

// 贪心
// 其实就是简化动态规划
func canJump(nums []int) bool {
	maxPath := 0
	for i, num := range nums {
		// 大于maxPath证明，i点不可达
		if i > maxPath {
			break
		}
		// 可达的话，算出当前点能到达的最远处
		if i+num > maxPath {
			maxPath = i + num
		}
	}
	return maxPath >= len(nums)-1
}
