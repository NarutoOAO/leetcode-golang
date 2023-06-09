package _300

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans := 0
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j] > dp[i] {
					dp[i] = dp[j]
				}
			}
		}
		dp[i]++
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
