package 周赛2

import "math"

func continuousSubarrays(nums []int) int64 {
	var res int64 = 0
	l := 0
	m := make(map[int]int)
	for r, v := range nums {
		m[v]++
		for max(m)-min(m) > 2 {
			v2 := nums[l]
			m[v2] -= 1
			if m[v2] == 0 {
				delete(m, v2)
			}
			l += 1
		}
		res += int64(r - l + 1)
	}
	return res
}

func max(m map[int]int) int {
	maxn := math.MinInt64
	for v, _ := range m {
		if v > maxn {
			maxn = v
		}
	}
	return maxn
}

func min(m map[int]int) int {
	minn := math.MaxInt64
	for v, _ := range m {
		if v < minn {
			minn = v
		}
	}
	return minn
}
