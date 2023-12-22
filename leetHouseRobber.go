package leetHouseRobber

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	var i int
	for i = 0; i == 0 && i < n; i++ {
		dp[i] = nums[i]
	}
	for i = 1; i == 1 && i < n; i++ {
		dp[i] = max(nums[i], nums[i-1])
	}
	for i = 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
