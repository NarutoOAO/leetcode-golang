package 周赛3

func maximumJumps(nums []int, target int) int {
	i, n := 0, len(nums)
	dp := make([]int, n)
	for i = 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if abs(nums[i]-nums[j]) <= target {
				if j != 0 && dp[j] == 0 {
					continue
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	if dp[n-1] == 0 {
		return -1
	}
	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
