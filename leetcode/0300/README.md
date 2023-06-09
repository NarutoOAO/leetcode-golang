最长递增子序列
DP j < i nums[j] < nums[i] dp[i] = max(dp[i], dp[j]) dp[i]++ 更新ans