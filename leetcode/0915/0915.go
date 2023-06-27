package _915

func partitionDisjoint(nums []int) int {
	curMax, leftMax, position, n := nums[0], nums[0], 0, len(nums)
	for i := 1; i < n-1; i++ {
		curMax = max(nums[i], curMax)
		if nums[i] < leftMax {
			leftMax = curMax
			position = i
		}
	}
	return position + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
