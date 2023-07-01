package _59

// 滑动窗口 使用队列记录最大值，每次超过边界把队列第一位弹出，每次加数字之前弹出之前小的数，从完成1-k次后往答案中加入数字
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	q := make([]int, 0)
	i, j := 0, 1-k
	for i < len(nums) {
		if j > 0 && nums[j-1] == q[0] {
			q = q[1:]
		}
		for len(q) > 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, nums[i])
		if j >= 0 {
			res[j] = q[0]
		}
		i++
		j++
	}
	return res
}
