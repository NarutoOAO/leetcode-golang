package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.OnesCount(9))
}

func numberOfGoodSubarraySplits(nums []int) int {
	i, j := 0, 0
	dp := make([]int, len(nums)+1)
	for i < len(nums) {
		if nums[i] == 1 {
			j = i
			dp[i] = 1
			break
		}
		i++
	}
	j++
	for j < len(nums) {
		if nums[j] == 1 {
			dp[j] = dp[i] + j - i - 1
			i = j
		} else {
			dp[j] = dp[i]
		}
		j++
	}
	return dp[len(nums)]
}
